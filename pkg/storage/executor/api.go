package executor

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/executor/internal/postgres"
)

const (
	postgresDialect = "postgres"
)

// New TODO issue#docs
func New(dialect string) *Executor {
	exec := &Executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// UserManager TODO issue#docs
type UserManager interface {
	AccessToken(domain.Token) (*repository.Token, error)
}

// Executor TODO issue#docs
type Executor struct {
	dialect string
	factory struct {
		NewUserManager func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO issue#docs
func (e *Executor) Dialect() string {
	return e.dialect
}

// UserManager TODO issue#docs
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
