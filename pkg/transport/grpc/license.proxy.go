package grpc

// Proxy TODO issue#docs
type Proxy interface {
	// Convert TODO issue#docs
	Convert() interface{}
}

// RegisterLicenseRequestProxy TODO issue#docs
type RegisterLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy RegisterLicenseRequestProxy) Convert() interface{} {
	return &RegisterLicenseRequest{}
}

// ExtendLicenseRequestProxy TODO issue#docs
type ExtendLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy ExtendLicenseRequestProxy) Convert() interface{} {
	return &ExtendLicenseRequest{}
}

// CheckLicenseRequestProxy TODO issue#docs
type CheckLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy CheckLicenseRequestProxy) Convert() interface{} {
	return &CheckLicenseRequest{}
}
