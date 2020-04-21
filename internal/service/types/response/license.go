package response

// CheckLicense TODO issue#docs
type CheckLicense struct {
	err error
}

// Cause TODO issue#docs
func (resp *CheckLicense) Cause() error {
	return resp.err
}

// Error TODO issue#docs
func (resp *CheckLicense) Error() string {
	if resp.err != nil {
		return resp.err.Error()
	}
	return ""
}

// HasError TODO issue#docs
func (resp *CheckLicense) HasError() bool {
	return resp.err != nil
}

// With TODO issue#docs
func (resp CheckLicense) With(err error) CheckLicense {
	resp.err = err
	return resp
}
