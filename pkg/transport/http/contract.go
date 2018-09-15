package http

import "github.com/kamilsk/guard/pkg/service/types"

// Service TODO issue#docs
type Service interface {
	// CheckLicense TODO issue#docs
	CheckLicense(types.License) error
}
