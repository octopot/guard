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
		License   types.ID `json:"license"`
		User      types.ID `json:"user"`
		Workplace types.ID `json:"workplace"`
		Metadata  metadata
	}

	request := scope{
		License:   types.ID(req.Header.Get("X-License")),
		User:      types.ID(req.Header.Get("X-User")),
		Workplace: types.ID(req.Header.Get("X-Workplace")),
		Metadata: metadata{
			Forward: req.Header.Get("X-Forwarded-For"),
			ID:      types.ID(req.Header.Get("X-Request-ID")),
			IP:      req.Header.Get("X-Real-IP"),
			URI:     req.Header.Get("X-Original-URI"),
		},
	}
	log.Println(request) // TODO issue#7

	license := types.License{Number: request.License, User: request.User, Workplace: request.Workplace}
	// TODO issue#6
	if err := srv.service.CheckLicense(license); err != nil {
		// TODO http.StatusUnauthorized and http.StatusForbidden must be have different logic
		http.Error(rw, http.StatusText(http.StatusPaymentRequired), http.StatusForbidden)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
