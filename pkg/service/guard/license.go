package guard

import (
	"errors"
	"math/rand"
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

type licenseManager struct {
	storage Storage
}

// Check TODO issue#docs
func (service *licenseManager) Check(license domain.License) error {
	if rand.New(rand.NewSource(time.Now().Unix())).Intn(5) > 2 {
		return errors.New("stub")
	}
	return nil
}
