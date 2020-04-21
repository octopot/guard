package guard

import (
	"context"

	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/guard/internal/service/types/request"
	"go.octolab.org/ecosystem/guard/internal/service/types/response"
)

type maintenanceService struct {
	storage accountStorage
}

// Install TODO issue#docs
func (service *maintenanceService) Install(ctx context.Context, req request.Install) response.Install {
	var (
		resp response.Install
		err  error
	)

	// TODO issue#6
	if req.Account == nil {
		return resp.With(errors.New("account is not provided"))
	}
	if len(req.Account.Users()) == 0 {
		return resp.With(errors.New("users are not provided"))
	}
	for _, user := range req.Account.Users() {
		if len(user.Tokens()) == 0 {
			user.WithDefaultToken()
		}
	}

	resp.Account, err = service.storage.RegisterAccount(ctx, req.Account)
	if err != nil {
		// TODO issue#6
		return resp.With(errors.WithMessage(err, "trying to do installation"))
	}
	return resp
}
