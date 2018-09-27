package guard

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

type licenseManager struct {
	storage Storage
}

// Check TODO issue#docs
func (service *licenseManager) Check(request request.License) (response response.License) {
	if rand.New(rand.NewSource(time.Now().Unix())).Intn(5) > 2 {
		// TODO issue#6
		// TODO issue#35
		return response.With(errors.New(http.StatusText(http.StatusPaymentRequired)))
	}
	return
}
