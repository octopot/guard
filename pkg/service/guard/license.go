package guard

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

type licenseService struct {
	disabled bool
	storage  Storage
}

// Check TODO issue#docs
func (service *licenseService) Check(ctx context.Context, req request.License) (resp response.License) {
	if service.disabled {
		// TODO issue#59 only logging request problems
		return
	}
	if rand.New(rand.NewSource(time.Now().Unix())).Intn(5) > 2 {
		// TODO issue#6
		// TODO issue#35
		return resp.With(errors.New(http.StatusText(http.StatusPaymentRequired)))
	}
	return
}
