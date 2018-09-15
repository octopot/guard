package types

import "github.com/kamilsk/go-kit/pkg/service/types"

// License TODO issue#docs
type License struct {
	Number    types.ID
	User      types.ID
	Workplace types.ID
}
