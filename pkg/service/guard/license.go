package guard

import (
	"context"
	"net/http"
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/service/guard/internal"
	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
	"github.com/pkg/errors"
)

type licenseService struct {
	disabled bool
	storage  Storage
}

// Check TODO issue#docs
func (service *licenseService) Check(ctx context.Context, req request.CheckLicense) (resp response.CheckLicense) {
	defer func() {
		if service.disabled && resp.HasError() {
			// TODO issue#59 only logging request problems
			resp = resp.With(nil)
		}
	}()

	var license domain.License

	switch {
	case !(req.ID.IsValid() && req.Employee.IsValid()) || !req.Workplace.IsValid():
		return resp.With(errors.New(http.StatusText(http.StatusBadRequest)))
	case req.ID.IsValid():
		entity, err := service.storage.LicenseByID(ctx, req.ID)
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

	checkers := []func(*domain.License) error{
		service.checkLifetimeLimits,
		service.checkRateLimits,
		service.checkRequestLimits,
		service.checkWorkplaceLimits,
	}
	results := make([]error, len(checkers))
	for i, check := range checkers {
		results[i] = check(&license)
	}

	return
}

func (service *licenseService) checkLifetimeLimits(license *domain.License) error {
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

func (service *licenseService) checkRateLimits(license *domain.License) error {
	if license.Rate.IsValid() {
		// TODO issue#future
		// errors.New(http.StatusText(http.StatusTooManyRequests))
		return nil
	}
	return nil
}

func (service *licenseService) checkRequestLimits(license *domain.License) error {
	if license.Requests > 0 && license.Requests < internal.LicenseRequests.IncrementFor(license.ID) {
		return errors.New(http.StatusText(http.StatusTooManyRequests))
	}
	return nil
}

func (service *licenseService) checkWorkplaceLimits(license *domain.License) error {
	if license.Workplaces > 0 {
		// errors.New(http.StatusText(http.StatusPaymentRequired))
		return nil
	}
	return nil
}
