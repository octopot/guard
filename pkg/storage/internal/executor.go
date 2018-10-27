package internal

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/internal/postgres"
	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/storage/types"
)

const (
	mysqlDialect    = "mysql"
	postgresDialect = "postgres"
)

// Constructor TODO issue#docs
type Constructor func(dialect string) Executor

// Executor TODO issue#docs
type Executor interface {
	// Dialect TODO issue#docs
	Dialect() string
	// LicenseManager TODO issue#docs
	LicenseManager(context.Context, *sql.Conn) LicenseManager
	// LicenseReader TODO issue#docs
	LicenseReader(context.Context, *sql.Conn) LicenseReader
	// UserManager TODO issue#docs
	UserManager(context.Context, *sql.Conn) UserManager
}

// New TODO issue#docs
func New(dialect string) Executor {
	exec := &executor{dialect: dialect}
	switch exec.dialect {
	case postgresDialect:
		exec.factory.NewLicenseManager = func(ctx context.Context, conn *sql.Conn) LicenseManager {
			return postgres.NewLicenseContext(ctx, conn)
		}
		exec.factory.NewLicenseReader = func(ctx context.Context, conn *sql.Conn) LicenseReader {
			return postgres.NewLicenseContext(ctx, conn)
		}
		exec.factory.NewUserManager = func(ctx context.Context, conn *sql.Conn) UserManager {
			return postgres.NewUserContext(ctx, conn)
		}
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
	Create(*types.Token, query.CreateLicense) (types.License, error)
	// Read TODO issue#docs
	Read(*types.Token, query.ReadLicense) (types.License, error)
	// Update TODO issue#docs
	Update(*types.Token, query.UpdateLicense) (types.License, error)
	// Delete TODO issue#docs
	Delete(*types.Token, query.DeleteLicense) (types.License, error)
	// Restore TODO issue#docs
	Restore(*types.Token, query.RestoreLicense) (types.License, error)

	// Draft TODO issue#docs

	// AddEmployee TODO issue#docs
	AddEmployee(*types.Token, query.LicenseEmployee) error
	// DeleteEmployee TODO issue#docs
	DeleteEmployee(*types.Token, query.LicenseEmployee) error
	// AddWorkplace TODO issue#docs
	AddWorkplace(*types.Token, query.LicenseWorkplace) error
	// DeleteWorkplace TODO issue#docs
	DeleteWorkplace(*types.Token, query.LicenseWorkplace) error
	// PushWorkplace TODO issue#docs
	PushWorkplace(*types.Token, query.LicenseWorkplace) error

	// issue#draft }
}

// LicenseReader TODO issue#docs
type LicenseReader interface {
	// GetByID TODO issue#docs
	GetByID(query.GetLicenseWithID) (types.License, error)
	// GetByEmployee TODO issue#docs
	GetByEmployee(query.GetEmployeeLicense) (types.License, error)
}

// UserManager TODO issue#docs
type UserManager interface {
	// AccessToken TODO issue#docs
	AccessToken(domain.Token) (*types.Token, error)
	// RegisterAccount TODO issue#docs
	RegisterAccount(query.RegisterAccount) (*types.Account, error)
	// RegisterUser TODO issue#docs
	RegisterUser(query.RegisterUser) (*types.User, error)
	// RegisterToken TODO issue#docs
	RegisterToken(query.RegisterToken) (*types.Token, error)
}

type executor struct {
	dialect string
	factory struct {
		NewLicenseManager func(context.Context, *sql.Conn) LicenseManager
		NewLicenseReader  func(context.Context, *sql.Conn) LicenseReader
		NewUserManager    func(context.Context, *sql.Conn) UserManager
	}
}

// Dialect TODO issue#docs
func (executor *executor) Dialect() string {
	return executor.dialect
}

// LicenseManager TODO issue#docs
func (executor *executor) LicenseManager(ctx context.Context, conn *sql.Conn) LicenseManager {
	return executor.factory.NewLicenseManager(ctx, conn)
}

// LicenseReader TODO issue#docs
func (executor *executor) LicenseReader(ctx context.Context, conn *sql.Conn) LicenseReader {
	return executor.factory.NewLicenseReader(ctx, conn)
}

// UserManager TODO issue#docs
func (executor *executor) UserManager(ctx context.Context, conn *sql.Conn) UserManager {
	return executor.factory.NewUserManager(ctx, conn)
}
