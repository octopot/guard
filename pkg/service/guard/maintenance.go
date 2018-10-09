package guard

import (
	"context"

	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
	"github.com/pkg/errors"
)

type maintenanceService struct {
	storage Storage
}

// Install TODO issue#docs
func (service *maintenanceService) Install(ctx context.Context, req request.Install) response.Install {
	var (
		resp response.Install
		err  error
	)
	resp.Account, err = service.storage.RegisterAccount(ctx, req.Account)
	if err != nil {
		// TODO issue#6
		return resp.With(errors.WithMessage(err, "trying to do installation"))
	}
	return resp
}
