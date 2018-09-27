package http

import "github.com/kamilsk/guard/pkg/service/request"

// Service TODO issue#docs
type Service interface {
	// CheckLicense TODO issue#docs
	CheckLicense(request.License) error
}
