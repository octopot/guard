package postgres

import (
	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/storage/types"
	"github.com/pkg/errors"
)

// TODO issue#draft {

// AddEmployee TODO issue#docs
func (scope licenseManager) AddEmployee(token *types.Token, data query.LicenseEmployee) error {
	license, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return readErr
	}
	q := `INSERT INTO "license_employee" ("license", "employee")
	      VALUES ($1, $2)
	      ON CONFLICT DO NOTHING`
	if _, execErr := scope.conn.ExecContext(scope.ctx, q, license.ID, data.Employee); execErr != nil {
		return errors.Wrapf(execErr,
			"user %q of account %q with token %q tried to add employee %q to the license %q",
			token.UserID, token.User.AccountID, token.ID, data.Employee, license.ID)
	}
	return nil
}

// DeleteEmployee TODO issue#docs
func (scope licenseManager) DeleteEmployee(token *types.Token, data query.LicenseEmployee) error {
	license, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return readErr
	}
	q := `DELETE FROM "license_employee"
	       WHERE "license" = $1 AND "employee" = $2`
	if _, execErr := scope.conn.ExecContext(scope.ctx, q, license.ID, data.Employee); execErr != nil {
		return errors.Wrapf(execErr,
			"user %q of account %q with token %q tried to delete employee %q from the license %q",
			token.UserID, token.User.AccountID, token.ID, data.Employee, license.ID)
	}
	return nil
}

// AddWorkplace TODO issue#docs
func (scope licenseManager) AddWorkplace(token *types.Token, data query.LicenseWorkplace) error {
	license, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return readErr
	}
	q := `INSERT INTO "license_workplace" ("license", "workplace")
	      VALUES ($1, $2)
	      ON CONFLICT DO NOTHING`
	if _, execErr := scope.conn.ExecContext(scope.ctx, q, license.ID, data.Workplace); execErr != nil {
		return errors.Wrapf(execErr,
			"user %q of account %q with token %q tried to add workplace %q to the license %q",
			token.UserID, token.User.AccountID, token.ID, data.Workplace, license.ID)
	}
	return nil
}

// DeleteWorkplace TODO issue#docs
func (scope licenseManager) DeleteWorkplace(token *types.Token, data query.LicenseWorkplace) error {
	license, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return readErr
	}
	q := `DELETE FROM "license_workplace"
	       WHERE "license" = $1 AND "workplace" = $2`
	if _, execErr := scope.conn.ExecContext(scope.ctx, q, license.ID, data.Workplace); execErr != nil {
		return errors.Wrapf(execErr,
			"user %q of account %q with token %q tried to delete workplace %q from the license %q",
			token.UserID, token.User.AccountID, token.ID, data.Workplace, license.ID)
	}
	return nil
}

// PushWorkplace TODO issue#docs
func (scope licenseManager) PushWorkplace(token *types.Token, data query.LicenseWorkplace) error {
	license, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return readErr
	}
	q := `UPDATE "license_workplace"
	         SET "updated_at" = now()
	       WHERE "license" = $1 AND "workplace" = $2`
	if _, execErr := scope.conn.ExecContext(scope.ctx, q, license.ID, data.Workplace); execErr != nil {
		return errors.Wrapf(execErr,
			"user %q of account %q with token %q tried to push workplace %q of the license %q",
			token.UserID, token.User.AccountID, token.ID, data.Workplace, license.ID)
	}
	return nil
}

// issue#draft }
