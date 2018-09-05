package router

import "net/http"

// API TODO issue#docs
type API interface {
	V1
}

// V1 TODO issue#docs
type V1 interface {
	// RegisterLicenseV1 TODO issue#docs
	RegisterLicenseV1(http.ResponseWriter, *http.Request)
	// ExtendLicenseV1 TODO issue#docs
	ExtendLicenseV1(http.ResponseWriter, *http.Request)
	// CheckLicenseV1 TODO issue#docs
	CheckLicenseV1(http.ResponseWriter, *http.Request)
}
