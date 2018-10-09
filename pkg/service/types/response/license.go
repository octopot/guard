package response

// License TODO issue#docs
type License struct {
	err error
}

// Cause TODO issue#docs
func (resp *License) Cause() error {
	return resp.err
}

// Error TODO issue#docs
func (resp *License) Error() string {
	if resp.err != nil {
		return resp.err.Error()
	}
	return ""
}

// HasError TODO issue#docs
func (resp *License) HasError() bool {
	return resp.err != nil
}

// With TODO issue#docs
func (resp License) With(err error) License {
	resp.err = err
	return resp
}
