package guard

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.octolab.org/ecosystem/guard/internal/platform/logger"
	"go.octolab.org/ecosystem/guard/internal/service/guard/internal"
	domain "go.octolab.org/ecosystem/guard/internal/service/types"
	"go.octolab.org/ecosystem/guard/internal/service/types/request"
	"go.octolab.org/ecosystem/guard/internal/service/types/response"
)

type licenseService struct {
	disabled bool
	storage  licenseStorage
}

// Check TODO issue#docs
func (service *licenseService) Check(ctx context.Context, req request.CheckLicense) (resp response.CheckLicense) {
	defer func() {
		if resp.HasError() {
			// TODO issue#59 only logging request problems
			l := logger.Default.With(zap.String("reason", resp.Error()))
			if !req.License.IsEmpty() {
				l = l.With(zap.String("license", req.License.String()))
			}
			if !req.Employee.IsEmpty() {
				l = l.With(zap.String("employee", req.Employee.String()))
			}
			if !req.Workplace.IsEmpty() {
				l = l.With(zap.String("workplace", req.Workplace.String()))
			}
			l.Error("license check")
			if service.disabled {
				resp = resp.With(nil)
			}
			_ = l.Sync()
		}
	}()

	var license domain.License

	switch {
	case !(req.License.IsValid() || req.Employee.IsValid()) || !req.Workplace.IsValid():
		return resp.With(errors.New(http.StatusText(http.StatusBadRequest)))
	case req.License.IsValid():
		entity, err := service.storage.LicenseByID(ctx, req.License)
		if err != nil {
			return resp.With(err)
		}
		license.ID, license.Contract = entity.ID, entity.Contract
	case req.Employee.IsValid():
		entity, err := service.storage.LicenseByEmployee(ctx, req.Employee)
		if err != nil {
			return resp.With(err)
		}
		license.ID, license.Contract = entity.ID, entity.Contract
	}

	// TODO issue#composite
	if err := service.checkLifetimeLimits(license, req.Workplace); err != nil {
		return resp.With(err)
	}
	if err := service.checkRateLimits(license, req.Workplace); err != nil {
		return resp.With(err)
	}
	if err := service.checkRequestLimits(license, req.Workplace); err != nil {
		return resp.With(err)
	}
	if err := service.checkWorkplaceLimits(license, req.Workplace); err != nil {
		return resp.With(err)
	}

	return
}

func (service *licenseService) checkLifetimeLimits(license domain.License, _ domain.ID) error {
	now := time.Now()
	if license.Since != nil {
		if license.Since.After(now) {
			return errors.New(http.StatusText(http.StatusUnprocessableEntity))
		}
	}
	if license.Until != nil {
		if license.Until.Before(now) {
			return errors.New(http.StatusText(http.StatusPaymentRequired))
		}
	}
	return nil
}

func (service *licenseService) checkRateLimits(license domain.License, _ domain.ID) error {
	if license.Rate.IsValid() {
		// TODO issue#future
		// errors.New(http.StatusText(http.StatusTooManyRequests))
		return nil
	}
	return nil
}

func (service *licenseService) checkRequestLimits(license domain.License, _ domain.ID) error {
	counter := internal.LicenseRequests
	if license.Requests > 0 && license.Requests < counter.Increment(license.ID) {
		go counter.Rollback(license.ID)
		return errors.New(http.StatusText(http.StatusTooManyRequests))
	}
	return nil
}

func (service *licenseService) checkWorkplaceLimits(license domain.License, workplace domain.ID) error {
	if !internal.LicenseWorkplaces.Acquire(license.ID, workplace, int(license.Workplaces)) {
		return errors.New(http.StatusText(http.StatusPaymentRequired))
	}
	return nil
}
