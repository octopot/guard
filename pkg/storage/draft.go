package storage

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// TODO issue#draft {

// AddEmployee TODO issue#docs
func (storage *Storage) AddEmployee(ctx context.Context, id domain.Token, data query.LicenseEmployee) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.Draft(ctx, conn).AddEmployee(token, data)
}

// DeleteEmployee TODO issue#docs
func (storage *Storage) DeleteEmployee(ctx context.Context, id domain.Token, data query.LicenseEmployee) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.Draft(ctx, conn).DeleteEmployee(token, data)
}

// AddWorkplace TODO issue#docs
func (storage *Storage) AddWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.Draft(ctx, conn).AddWorkplace(token, data)
}

// DeleteWorkplace TODO issue#docs
func (storage *Storage) DeleteWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.Draft(ctx, conn).DeleteWorkplace(token, data)
}

// PushWorkplace TODO issue#docs
func (storage *Storage) PushWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.Draft(ctx, conn).PushWorkplace(token, data)
}

// issue#draft }
