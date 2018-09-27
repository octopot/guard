package http

import (
	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

// Service TODO issue#docs
type Service interface {
	// CheckLicense TODO issue#docs
	CheckLicense(request.License) response.License
}
