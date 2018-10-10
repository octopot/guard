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
	mysqlDialect    = "mysql"
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

		// TODO issue#draft {

		exec.factory.NewDraft = func(ctx context.Context, conn *sql.Conn) Draft {
			return postgres.NewLicenseContext(ctx, conn)
		}

		// issue#draft }

	case mysqlDialect:
		fallthrough
	default:
		panic(fmt.Sprintf("not supported dialect %q is provided", exec.dialect))
	}
	return exec
}

// LicenseManager TODO issue#docs
type LicenseManager interface {
	// Create TODO issue#docs
	Create(*repository.Token, query.CreateLicense) (repository.License, error)
	// Read TODO issue#docs
	Read(*repository.Token, query.ReadLicense) (repository.License, error)
	// Update TODO issue#docs
	Update(*repository.Token, query.UpdateLicense) (repository.License, error)
	// Delete TODO issue#docs
	Delete(*repository.Token, query.DeleteLicense) (repository.License, error)
	// Restore TODO issue#docs
	Restore(*repository.Token, query.RestoreLicense) (repository.License, error)
}

// UserManager TODO issue#docs
type UserManager interface {
	// AccessToken TODO issue#docs
	AccessToken(domain.Token) (*repository.Token, error)
	// RegisterAccount TODO issue#docs
	RegisterAccount(query.RegisterAccount) (*repository.Account, error)
	// RegisterUser TODO issue#docs
	RegisterUser(query.RegisterUser) (*repository.User, error)
	// RegisterToken TODO issue#docs
	RegisterToken(query.RegisterToken) (*repository.Token, error)
}

// Executor TODO issue#docs
type Executor struct {
	dialect string
	factory struct {
		NewLicenseManager func(context.Context, *sql.Conn) LicenseManager
		NewUserManager    func(context.Context, *sql.Conn) UserManager

		// TODO issue#draft {

		NewDraft func(context.Context, *sql.Conn) Draft

		// issue#draft }
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

// TODO issue#draft {

// Draft TODO issue#docs
func (e *Executor) Draft(ctx context.Context, conn *sql.Conn) Draft {
	return e.factory.NewDraft(ctx, conn)
}

// Draft TODO issue#docs
type Draft interface {
	// AddEmployee TODO issue#docs
	AddEmployee(*repository.Token, query.LicenseEmployee) error
	// DeleteEmployee TODO issue#docs
	DeleteEmployee(*repository.Token, query.LicenseEmployee) error
	// AddWorkplace TODO issue#docs
	AddWorkplace(*repository.Token, query.LicenseWorkplace) error
	// DeleteWorkplace TODO issue#docs
	DeleteWorkplace(*repository.Token, query.LicenseWorkplace) error
	// PushWorkplace TODO issue#docs
	PushWorkplace(*repository.Token, query.LicenseWorkplace) error
}

// issue#draft }
