package http

import (
	"net/http"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/service/request"
)

// CheckLicenseV1 TODO issue#docs
func (srv *server) CheckLicenseV1(rw http.ResponseWriter, req *http.Request) {
	if err := srv.service.CheckLicense(request.License{
		Number:    domain.ID(req.Header.Get("X-License")),
		Employee:  domain.ID(req.Header.Get("X-Employee")),
		Workplace: domain.ID(req.Header.Get("X-Workplace")),
		Metadata: request.Metadata{
			Forward: req.Header.Get("X-Forwarded-For"),
			ID:      domain.ID(req.Header.Get("X-Request-ID")),
			IP:      req.Header.Get("X-Real-IP"),
			URI:     req.Header.Get("X-Original-URI"),
		},
	}); err != nil {
		// TODO issue#6
		// TODO issue#34
		// TODO issue#35
		rw.Header().Set("X-Reason", http.StatusText(http.StatusPaymentRequired))
		http.Error(rw, http.StatusText(http.StatusPaymentRequired), http.StatusForbidden)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
