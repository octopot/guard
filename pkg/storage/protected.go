package storage

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// ExtendLicense TODO issue#docs
func (storage *Storage) ExtendLicense(ctx context.Context, tokenID domain.ID) error {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return err
	}
	defer closer()

	_, err = storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return err
	}

	return nil
}

// RegisterLicense TODO issue#docs
func (storage *Storage) RegisterLicense(ctx context.Context, tokenID domain.ID) error {
	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return err
	}
	defer closer()

	_, err = storage.exec.UserManager(ctx, conn).Token(tokenID)
	if err != nil {
		return err
	}

	return nil
}
