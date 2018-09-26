package storage

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// ExtendLicense TODO issue#docs
func (storage *Storage) ExtendLicense(ctx context.Context, id domain.Token, data query.ExtendLicense) (repository.License, error) {
	var license repository.License

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return license, err
	}
	defer closer()

	_, err = storage.exec.UserManager(ctx, conn).AccessToken(id)
	if err != nil {
		return license, err
	}

	return license, nil
}

// ReadLicense TODO issue#docs
func (storage *Storage) ReadLicense(ctx context.Context, id domain.Token, data query.ReadLicense) (repository.License, error) {
	var license repository.License

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return license, err
	}
	defer closer()

	_, err = storage.exec.UserManager(ctx, conn).AccessToken(id)
	if err != nil {
		return license, err
	}

	return license, nil
}

// RegisterLicense TODO issue#docs
func (storage *Storage) RegisterLicense(ctx context.Context, id domain.Token, data query.RegisterLicense) (repository.License, error) {
	var license repository.License

	conn, closer, err := storage.connection(ctx)
	if err != nil {
		return license, err
	}
	defer closer()

	_, err = storage.exec.UserManager(ctx, conn).AccessToken(id)
	if err != nil {
		return license, err
	}

	return license, nil
}
