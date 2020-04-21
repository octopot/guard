package request

import "go.octolab.org/ecosystem/guard/internal/storage/query"

// Install TODO issue#docs
type Install struct {
	Account *query.RegisterAccount
}
