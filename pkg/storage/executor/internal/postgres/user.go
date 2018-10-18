package postgres

import (
	"context"
	"database/sql"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/storage/types"
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
func (scope userManager) AccessToken(id domain.Token) (*types.Token, error) {
	var (
		account = types.Account{}
		token   = types.Token{ID: id}
		user    = types.User{}
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

// RegisterAccount TODO issue#docs
func (scope userManager) RegisterAccount(data query.RegisterAccount) (*types.Account, error) {
	entity := types.Account{Name: data.Name}
	q := `INSERT INTO "account" ("id", "name")
	      SELECT coalesce($1, uuid_generate_v4()), $2
	       WHERE NOT exists(SELECT "id" FROM "account")
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.Name)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return nil, errors.Wrap(scanErr, "trying to register a new account")
	}
	return &entity, nil
}

// RegisterUser TODO issue#docs
func (scope userManager) RegisterUser(data query.RegisterUser) (*types.User, error) {
	entity, account := types.User{Name: data.Name}, data.Account()
	if account == nil || account.ID == nil {
		return nil, errors.New("the account is not specified")
	}
	entity.AccountID = *account.ID
	q := `INSERT INTO "user" ("id", "account_id", "name")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.AccountID, entity.Name)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return nil, errors.Wrap(scanErr, "trying to register a new user")
	}
	return &entity, nil
}

// RegisterToken TODO issue#docs
func (scope userManager) RegisterToken(data query.RegisterToken) (*types.Token, error) {
	entity, user := types.Token{ExpiredAt: data.ExpiredAt}, data.User()
	if user == nil || user.ID == nil {
		return nil, errors.New("the user is not specified")
	}
	entity.UserID = *user.ID
	q := `INSERT INTO "token" ("id", "user_id", "expired_at")
	      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
	   RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, entity.UserID, entity.ExpiredAt)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return nil, errors.Wrap(scanErr, "trying to register a new token")
	}
	return &entity, nil
}
