package request

import "github.com/kamilsk/guard/pkg/storage/query"

// Install TODO issue#docs
type Install struct {
	Account *query.RegisterAccount
}
