package http

import (
	"net/http"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/service/types/request"
)

// CheckLicenseV1 TODO issue#docs
func (srv *server) CheckLicenseV1(rw http.ResponseWriter, req *http.Request) {
	if response := srv.service.CheckLicense(request.License{
		ID:        domain.ID(req.Header.Get("X-License")),
		Employee:  domain.ID(req.Header.Get("X-Employee")),
		Workplace: domain.ID(req.Header.Get("X-Workplace")),
		Metadata: request.Metadata{
			Forward: req.Header.Get("X-Forwarded-For"),
			ID:      domain.ID(req.Header.Get("X-Request-ID")),
			IP:      req.Header.Get("X-Real-IP"),
			URI:     req.Header.Get("X-Original-URI"),
		},
	}); response.HasError() {
		rw.Header().Set("X-Reason", response.Error())
		// TODO issue#34 http.{StatusUnauthorized StatusForbidden StatusInternalServerError StatusServiceUnavailable}
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
