package storage

import (
	"context"
	"database/sql"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/storage/types"
)

// LicenseByID TODO issue#docs
func (storage *Storage) LicenseByID(ctx context.Context, id domain.ID) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	return storage.exec.LicenseReader(ctx, conn).GetByID(query.GetLicenseWithID{ID: id})
}

// LicenseByEmployee TODO issue#docs
func (storage *Storage) LicenseByEmployee(ctx context.Context, employee domain.ID) (types.License, error) {
	var license types.License

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return license, connErr
	}
	defer closer()

	return storage.exec.LicenseReader(ctx, conn).GetByEmployee(query.GetEmployeeLicense{Employee: employee})
}

// RegisterAccount TODO issue#docs
func (storage *Storage) RegisterAccount(ctx context.Context, data *query.RegisterAccount) (*types.Account, error) {

	// register is possible if no account exists
	// it is checked on executor layer (postgres)

	conn, closer, connErr := storage.connection(ctx)
	if connErr != nil {
		return nil, connErr
	}
	defer closer()

	tx, txErr := conn.BeginTx(ctx, &sql.TxOptions{})
	if txErr != nil {
		return nil, txErr
	}

	var (
		account *types.Account
		err     error
	)
	func() {
		defer func() {
			finalize := tx.Commit
			if err != nil {
				finalize = tx.Rollback
			}
			_ = finalize() // TODO issue#composite
		}()

		manager := storage.exec.UserManager(ctx, conn)
		account, err = manager.RegisterAccount(*data)
		if err != nil {
			return
		}
		data.ID = &account.ID
		for _, regUser := range data.Users() {
			var user *types.User
			user, err = manager.RegisterUser(*regUser)
			if err != nil {
				return
			}
			regUser.ID = &user.ID
			user.Account, account.Users = account, append(account.Users, user)
			for _, regToken := range regUser.Tokens() {
				var token *types.Token
				token, err = manager.RegisterToken(*regToken)
				token.User, user.Tokens = user, append(user.Tokens, token)
			}
		}
	}()
	return account, err
}
