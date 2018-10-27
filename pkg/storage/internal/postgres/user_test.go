package postgres_test

import (
	"context"
	"errors"
	"testing"
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/internal"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	. "github.com/kamilsk/guard/pkg/storage/internal/postgres"
)

func TestUserManager(t *testing.T) {
	id, token := domain.ID("10000000-2000-4000-8000-160000000000"), domain.Token("10000000-2000-4000-8000-160000000000")
	t.Run("token", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`SELECT "(?:.+)" FROM "token"`).
				WithArgs(token).
				WillReturnRows(
					sqlmock.
						NewRows([]string{
							"user_id", "expired_at", "created_at",
							"account_id", "name", "created_at", "updated_at",
							"name", "created_at", "updated_at",
						}).
						AddRow(
							id, time.Now(), time.Now(),
							id, "test", time.Now(), time.Now(),
							"test", time.Now(), time.Now(),
						),
				)

			var exec internal.UserManager = NewUserContext(ctx, conn)
			tkn, err := exec.AccessToken(token)
			assert.NoError(t, err)
			assert.NotEmpty(t, tkn.UserID)
			assert.NotEmpty(t, tkn.ExpiredAt)
			assert.NotEmpty(t, tkn.CreatedAt)
			assert.NotEmpty(t, tkn.User)
			assert.NotEmpty(t, tkn.User.AccountID)
			assert.NotEmpty(t, tkn.User.Name)
			assert.NotEmpty(t, tkn.User.CreatedAt)
			assert.NotEmpty(t, tkn.User.UpdatedAt)
			assert.NotEmpty(t, tkn.User.Account)
			assert.NotEmpty(t, tkn.User.Account.Name)
			assert.NotEmpty(t, tkn.User.Account.CreatedAt)
			assert.NotEmpty(t, tkn.User.Account.UpdatedAt)
		})
		t.Run("database error", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			conn, err := db.Conn(ctx)
			assert.NoError(t, err)
			defer conn.Close()

			mock.
				ExpectQuery(`SELECT "(?:.+)" FROM "token"`).
				WithArgs(id).
				WillReturnError(errors.New("test"))

			var exec internal.UserManager = NewUserContext(ctx, conn)
			tkn, err := exec.AccessToken(token)
			assert.Error(t, err)
			assert.Nil(t, tkn)
		})
	})
}
