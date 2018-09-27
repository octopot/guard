package guard

import (
	"errors"
	"math/rand"
	"time"

	"github.com/kamilsk/guard/pkg/service/request"
)

type licenseManager struct {
	storage Storage
}

// Check TODO issue#docs
func (service *licenseManager) Check(license request.License) error {
	if rand.New(rand.NewSource(time.Now().Unix())).Intn(5) > 2 {
		return errors.New("stub")
	}
	return nil
}
