package router

import "net/http"

// API TODO issue#docs
type API interface {
	V1
}

// V1 TODO issue#docs
type V1 interface {
	// CheckLicenseV1 TODO issue#docs
	CheckLicenseV1(http.ResponseWriter, *http.Request)
}
