package http

import (
	"context"

	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

// Service TODO issue#docs
type Service interface {
	// CheckLicense TODO issue#docs
	CheckLicense(context.Context, request.CheckLicense) response.CheckLicense
}
