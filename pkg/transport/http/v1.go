package http

import (
	"log"
	"net/http"

	"github.com/kamilsk/go-kit/pkg/service/types"
)

// CheckLicenseV1 TODO issue#docs
func (srv *server) CheckLicenseV1(rw http.ResponseWriter, req *http.Request) {

	request := types.ID(req.Header.Get("X-Request-ID"))
	token := types.ID(req.Header.Get("X-Token"))
	udid := types.ID(req.Header.Get("X-UDID"))
	license := types.ID(req.Header.Get("X-License"))

	log.Println("request:", request, "token:", token, "udid:", udid, "license:", license)
	//rw.WriteHeader(http.StatusUnauthorized)
	rw.WriteHeader(http.StatusForbidden)
}
