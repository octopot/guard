package executor

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/executor/internal/postgres"
	"github.com/kamilsk/guard/pkg/storage/query"
)

const (
	postgresDialect = "postgres"
)

// New TODO issue#docs
func New(dialect string) *Executor {
	exec := &Executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewLicenseManager = func(ctx context.Context, conn *sql.Conn) LicenseManager {
			return postgres.NewLicenseContext(ctx, conn)
		}
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// LicenseManager TODO issue#docs
type LicenseManager interface {
	Create(*repository.Token, query.CreateLicense) (repository.License, error)
	Read(*repository.Token, query.ReadLicense) (repository.License, error)
	Update(*repository.Token, query.UpdateLicense) (repository.License, error)
	Delete(*repository.Token, query.DeleteLicense) (repository.License, error)
	Restore(*repository.Token, query.RestoreLicense) (repository.License, error)
}

// UserManager TODO issue#docs
type UserManager interface {
	AccessToken(domain.Token) (*repository.Token, error)
}

// Executor TODO issue#docs
type Executor struct {
	dialect string
	factory struct {
		NewLicenseManager func(context.Context, *sql.Conn) LicenseManager
		NewUserManager    func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO issue#docs
func (e *Executor) Dialect() string {
	return e.dialect
}

// LicenseManager TODO issue#docs
func (e *Executor) LicenseManager(ctx context.Context, conn *sql.Conn) LicenseManager {
	return e.factory.NewLicenseManager(ctx, conn)
}

// UserManager TODO issue#docs
func (e *Executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return e.factory.NewUserManager(ctx, conn)
}
