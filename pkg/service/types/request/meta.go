package request

import domain "github.com/kamilsk/guard/pkg/service/types"

// Metadata TODO issue#docs
type Metadata struct {
	Forward string
	ID      domain.ID
	IP      string
	URI     string
}
