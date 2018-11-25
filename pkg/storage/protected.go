package storage

import (
	"context"
	"database/sql"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/storage/types"
	"github.com/pkg/errors"
)

// RegisterLicense TODO issue#docs
func (storage *Storage) RegisterLicense(ctx context.Context, id domain.Token, data query.RegisterLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	defer func() { _ = tx.Rollback() }()

	manager := storage.exec.LicenseManager(ctx, conn)
	license, execErr := manager.Read(token, query.ReadLicense{ID: data.ID})
	if execErr == nil {
		license, execErr = manager.Update(token, query.UpdateLicense{ID: data.ID, Contract: data.Contract})
	} else if errors.Cause(execErr) == sql.ErrNoRows {
		license, execErr = manager.Create(token, query.CreateLicense{ID: &data.ID, Contract: data.Contract})
	}
	if execErr != nil {
		return license, execErr
	}

	return license, tx.Commit()
}

// CreateLicense TODO issue#docs
func (storage *Storage) CreateLicense(ctx context.Context, id domain.Token, data query.CreateLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	defer func() { _ = tx.Rollback() }()

	license, execErr := storage.exec.LicenseManager(ctx, conn).Create(token, data)
	if execErr != nil {
		return license, execErr
	}
	return license, tx.Commit()
}

// ReadLicense TODO issue#docs
func (storage *Storage) ReadLicense(ctx context.Context, id domain.Token, data query.ReadLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	return storage.exec.LicenseManager(ctx, conn).Read(token, data)
}

// UpdateLicense TODO issue#docs
func (storage *Storage) UpdateLicense(ctx context.Context, id domain.Token, data query.UpdateLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	defer func() { _ = tx.Rollback() }()

	license, execErr := storage.exec.LicenseManager(ctx, conn).Update(token, data)
	if execErr != nil {
		return license, execErr
	}
	return license, tx.Commit()
}

// DeleteLicense TODO issue#docs
func (storage *Storage) DeleteLicense(ctx context.Context, id domain.Token, data query.DeleteLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	defer func() { _ = tx.Rollback() }()

	license, execErr := storage.exec.LicenseManager(ctx, conn).Delete(token, data)
	if execErr != nil {
		return license, execErr
	}
	return license, tx.Commit()
}

// RestoreLicense TODO issue#docs
func (storage *Storage) RestoreLicense(ctx context.Context, id domain.Token, data query.RestoreLicense) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return license, authErr
	}

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return license, txErr
	}
	defer func() { _ = tx.Rollback() }()

	license, execErr := storage.exec.LicenseManager(ctx, conn).Restore(token, data)
	if execErr != nil {
		return license, execErr
	}
	return license, tx.Commit()
}

// TODO issue#draft {

// AddEmployee TODO issue#docs
func (storage *Storage) AddEmployee(ctx context.Context, id domain.Token, data query.LicenseEmployee) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.LicenseManager(ctx, conn).AddEmployee(token, data)
}

// DeleteEmployee TODO issue#docs
func (storage *Storage) DeleteEmployee(ctx context.Context, id domain.Token, data query.LicenseEmployee) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.LicenseManager(ctx, conn).DeleteEmployee(token, data)
}

// AddWorkplace TODO issue#docs
func (storage *Storage) AddWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.LicenseManager(ctx, conn).AddWorkplace(token, data)
}

// DeleteWorkplace TODO issue#docs
func (storage *Storage) DeleteWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.LicenseManager(ctx, conn).DeleteWorkplace(token, data)
}

// PushWorkplace TODO issue#docs
func (storage *Storage) PushWorkplace(ctx context.Context, id domain.Token, data query.LicenseWorkplace) error {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return authErr
	}

	return storage.exec.LicenseManager(ctx, conn).PushWorkplace(token, data)
}

// LicenseWorkplaces TODO issue#docs
func (storage *Storage) LicenseWorkplaces(ctx context.Context, id domain.Token, data query.WorkplaceList) (
	[]types.Workplace,
	error,
) {
	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return nil, connErr
	}
	defer func() { _ = closer() }()

	token, authErr := storage.exec.UserManager(ctx, conn).AccessToken(id)
	if authErr != nil {
		return nil, authErr
	}

	return storage.exec.LicenseManager(ctx, conn).Workplaces(token, data)
}

// issue#draft }
