package api

import (
	"context"

	"go.octolab.org/ecosystem/guard/internal/service/types/request"
	"go.octolab.org/ecosystem/guard/internal/service/types/response"
)

// Service TODO issue#docs
type Service interface {
	// CheckLicense TODO issue#docs
	CheckLicense(context.Context, request.CheckLicense) response.CheckLicense
}
