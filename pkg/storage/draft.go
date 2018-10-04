package storage

import (
	"context"

	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// RegisterAccount TODO issue#docs
func (storage Storage) RegisterAccount(ctx context.Context, data query.RegisterAccount) (repository.Account, error) {
	return repository.Account{}, nil
}
