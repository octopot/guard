package response

// License TODO issue#docs
type License struct {
	err error
}

// Cause TODO issue#docs
func (l *License) Cause() error {
	return l.err
}

// Error TODO issue#docs
func (l *License) Error() string {
	if l.err != nil {
		return l.err.Error()
	}
	return ""
}

// HasError TODO issue#docs
func (l *License) HasError() bool {
	return l.err != nil
}

// With TODO issue#docs
func (l License) With(err error) License {
	l.err = err
	return l
}
