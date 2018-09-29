package postgres

import (
	"context"
	"database/sql"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/pkg/errors"
)

// NewUserContext TODO issue#docs
func NewUserContext(ctx context.Context, conn *sql.Conn) userManager {
	return userManager{ctx, conn}
}

type userManager struct {
	ctx  context.Context
	conn *sql.Conn
}

// AccessToken TODO issue#docs
func (scope userManager) AccessToken(id domain.Token) (*repository.Token, error) {
	var (
		account = repository.Account{}
		token   = repository.Token{ID: id}
		user    = repository.User{}
	)
	q := `SELECT "t"."user_id", "t"."expired_at", "t"."created_at",
	             "u"."account_id", "u"."name", "u"."created_at", "u"."updated_at",
	             "a"."name", "a"."created_at", "a"."updated_at"
	        FROM "token" "t"
	  INNER JOIN "user" "u" ON "u"."id" = "t"."user_id"
	  INNER JOIN "account" "a" ON "a"."id" = "u"."account_id"
	       WHERE "t"."id" = $1 AND ("t"."expired_at" IS NULL OR "t"."expired_at" > now()) AND NOT "t"."revoked"
	         AND "u"."deleted_at" IS NULL
	         AND "a"."deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, token.ID)
	if scanErr := row.Scan(
		&token.UserID, &token.ExpiredAt, &token.CreatedAt,
		&user.AccountID, &user.Name, &user.CreatedAt, &user.UpdatedAt,
		&account.Name, &account.CreatedAt, &account.UpdatedAt,
	); scanErr != nil {
		return nil, errors.Wrapf(scanErr, "trying to get token by its id %q", id)
	}
	user.ID, account.ID = token.UserID, user.AccountID
	token.User, user.Account = &user, &account
	user.Tokens, account.Users = append(user.Tokens, &token), append(account.Users, &user)
	return &token, nil
}
