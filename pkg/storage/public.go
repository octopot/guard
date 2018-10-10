package storage

import (
	"context"
	"database/sql"

	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// RegisterAccount TODO issue#docs
func (storage Storage) RegisterAccount(ctx context.Context, data *query.RegisterAccount) (*repository.Account, error) {

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
		account *repository.Account
		err     error
	)
	func() {
		defer func() {
			finalize := tx.Commit
			if err != nil {
				finalize = tx.Rollback
			}
			err = finalize()
		}()

		manager := storage.exec.UserManager(ctx, conn)
		account, err = manager.RegisterAccount(*data)
		if err != nil {
			return
		}
		data.ID = &account.ID
		for _, regUser := range data.Users() {
			var user *repository.User
			user, err = manager.RegisterUser(*regUser)
			if err != nil {
				return
			}
			regUser.ID = &user.ID
			user.Account, account.Users = account, append(account.Users, user)
			for _, regToken := range regUser.Tokens() {
				var token *repository.Token
				token, err = manager.RegisterToken(*regToken)
				token.User, user.Tokens = user, append(user.Tokens, token)
			}
		}
	}()
	return account, err
}
