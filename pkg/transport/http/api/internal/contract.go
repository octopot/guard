package internal

import "net/http"

// API TODO issue#docs
type API interface {
	v1
}

type v1 interface {
	// CheckLicenseV1 TODO issue#docs
	CheckLicenseV1(http.ResponseWriter, *http.Request)
}
