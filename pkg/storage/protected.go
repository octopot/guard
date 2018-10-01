package storage

import (
	"context"
	"database/sql"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// CreateLicense TODO issue#docs
func (storage *Storage) CreateLicense(ctx context.Context, id domain.Token, data query.CreateLicense) (repository.License, error) {
	var license repository.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	license, execErr := storage.exec.LicenseManager(ctx, conn).Create(token, data)
	if execErr != nil {
		_ = tx.Rollback() // TODO issue#composite
		return license, execErr
	}
	return license, tx.Commit()
}

// ReadLicense TODO issue#docs
func (storage *Storage) ReadLicense(ctx context.Context, id domain.Token, data query.ReadLicense) (repository.License, error) {
	var license repository.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	return storage.exec.LicenseManager(ctx, conn).Read(token, data)
}

// UpdateLicense TODO issue#docs
func (storage *Storage) UpdateLicense(ctx context.Context, id domain.Token, data query.UpdateLicense) (repository.License, error) {
	var license repository.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	license, execErr := storage.exec.LicenseManager(ctx, conn).Update(token, data)
	if execErr != nil {
		_ = tx.Rollback() // TODO issue#composite
		return license, execErr
	}
	return license, tx.Commit()
}

// DeleteLicense TODO issue#docs
func (storage *Storage) DeleteLicense(ctx context.Context, id domain.Token, data query.DeleteLicense) (repository.License, error) {
	var license repository.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	license, execErr := storage.exec.LicenseManager(ctx, conn).Delete(token, data)
	if execErr != nil {
		_ = tx.Rollback() // TODO issue#composite
		return license, execErr
	}
	return license, tx.Commit()
}

// RestoreLicense TODO issue#docs
func (storage *Storage) RestoreLicense(ctx context.Context, id domain.Token, data query.RestoreLicense) (repository.License, error) {
	var license repository.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	license, execErr := storage.exec.LicenseManager(ctx, conn).Restore(token, data)
	if execErr != nil {
		_ = tx.Rollback() // TODO issue#composite
		return license, execErr
	}
	return license, tx.Commit()
}

// ---

// RegisterLicense TODO issue#docs
func (storage *Storage) RegisterLicense(ctx context.Context, id domain.Token, data query.RegisterLicense) (repository.License, error) {
	return storage.CreateLicense(ctx, id, query.CreateLicense{ID: &data.ID, Contract: data.Contract})
}
