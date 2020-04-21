package api

import (
	"net/http"

	domain "go.octolab.org/ecosystem/guard/internal/service/types"
	"go.octolab.org/ecosystem/guard/internal/service/types/request"
)

// CheckLicenseV1 TODO issue#docs
func (server *server) CheckLicenseV1(rw http.ResponseWriter, req *http.Request) {
	metadata := domain.MetadataFromRequest(req)
	if response := server.service.CheckLicense(req.Context(), request.CheckLicense{
		License:   metadata.License(),
		Employee:  metadata.Employee(),
		Workplace: metadata.Workplace(),
	}); response.HasError() {
		rw.Header().Set("X-Reason", response.Error())
		// TODO issue#34 http.{StatusUnauthorized StatusForbidden StatusInternalServerError StatusServiceUnavailable}
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
