package grpc

// Proxy TODO issue#docs
type Proxy interface {
	// Convert TODO issue#docs
	Convert() interface{}
}

// ExtendLicenseRequestProxy TODO issue#docs
type ExtendLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy ExtendLicenseRequestProxy) Convert() interface{} {
	return &ExtendLicenseRequest{}
}

// ReadLicenseRequestProxy TODO issue#docs
type ReadLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy ReadLicenseRequestProxy) Convert() interface{} {
	return &ReadLicenseRequest{}
}

// RegisterLicenseRequestProxy TODO issue#docs
type RegisterLicenseRequestProxy struct{}

// Convert TODO issue#docs
func (proxy RegisterLicenseRequestProxy) Convert() interface{} {
	return &RegisterLicenseRequest{}
}
