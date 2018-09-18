package http

import (
	"log"
	"net/http"

	"github.com/kamilsk/guard/pkg/service/types"
)

// CheckLicenseV1 TODO issue#docs
func (srv *server) CheckLicenseV1(rw http.ResponseWriter, req *http.Request) {
	type metadata struct {
		Forward string   `json:"forward"`
		ID      types.ID `json:"id"`
		IP      string   `json:"ip"`
		URI     string   `json:"uri"`
	}
	type scope struct {
		License  types.ID `json:"license"`
		UDID     types.ID `json:"udid"`
		User     types.ID `json:"user"`
		Metadata metadata
	}

	request := scope{
		License: types.ID(req.Header.Get("X-License")),
		UDID:    types.ID(req.Header.Get("X-UDID")),
		User:    types.ID(req.Header.Get("X-User")),
		Metadata: metadata{
			Forward: req.Header.Get("X-Forwarded-For"),
			ID:      types.ID(req.Header.Get("X-Request-ID")),
			IP:      req.Header.Get("X-Real-IP"),
			URI:     req.Header.Get("X-Original-URI"),
		},
	}
	log.Println(request) // TODO issue#7

	license := types.License{Number: request.License, User: request.User, Workplace: request.UDID}
	// TODO issue#6
	if err := srv.service.CheckLicense(license); err != nil {
		// TODO http.StatusUnauthorized and http.StatusForbidden must be have different logic
		http.Error(rw, http.StatusText(http.StatusPaymentRequired), http.StatusForbidden)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
