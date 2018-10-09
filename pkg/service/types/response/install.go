package response

import repository "github.com/kamilsk/guard/pkg/storage/types"

// Install TODO issue#docs
type Install struct {
	Account *repository.Account

	err error
}

// Cause TODO issue#docs
func (l *Install) Cause() error {
	return l.err
}

// Error TODO issue#docs
func (l *Install) Error() string {
	if l.err != nil {
		return l.err.Error()
	}
	return ""
}

// HasError TODO issue#docs
func (l *Install) HasError() bool {
	return l.err != nil
}

// With TODO issue#docs
func (l Install) With(err error) Install {
	l.err = err
	return l
}
