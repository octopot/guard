package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// NewLicenseContext TODO issue#docs
func NewLicenseContext(ctx context.Context, conn *sql.Conn) licenseManager {
	return licenseManager{ctx, conn}
}

type licenseManager struct {
	ctx  context.Context
	conn *sql.Conn
}

// Create TODO issue#docs
func (scope licenseManager) Create(token *repository.Token, data query.CreateLicense) (repository.License, error) {
	entity := repository.License{Contract: data.Contract}
	before, encodeErr := json.Marshal(domain.Contract{})
	if encodeErr != nil {
		return entity, encodeErr
	}
	after, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, encodeErr
	}
	{
		q := `INSERT INTO "license" ("id", "account_id", "contract")
		      VALUES (coalesce($1, uuid_generate_v4()), $2, $3)
		   RETURNING "id", "created_at"`
		row := scope.conn.QueryRowContext(scope.ctx, q, data.ID, token.User.AccountID, after)
		if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
			return entity, scanErr
		}
	}
	{
		q := `INSERT INTO "license_audit" ("license_id", "contract", "what", "when", "who", "with")
		      VALUES ($1, $2, $3, $4, $5, $6)`
		if _, execErr := scope.conn.ExecContext(scope.ctx, q, entity.ID, before,
			repository.Create, entity.CreatedAt, token.UserID, token.ID); execErr != nil {
			return entity, execErr
		}
	}
	return entity, nil
}

// Read TODO issue#docs
func (scope licenseManager) Read(token *repository.Token, data query.ReadLicense) (repository.License, error) {
	entity, encoded := repository.License{ID: data.ID}, []byte(nil)
	q := `SELECT "contract", "created_at", "updated_at", "deleted_at"
	        FROM "license"
	       WHERE "id" = $1 AND "account_id" = $2`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, token.User.AccountID)
	if err := row.Scan(&encoded, &entity.CreatedAt, &entity.UpdatedAt, &entity.DeletedAt); err != nil {
		return entity, err
	}
	if err := json.Unmarshal(encoded, &entity.Contract); err != nil {
		return entity, err
	}
	return entity, nil
}

// Update TODO issue#docs
func (scope licenseManager) Update(token *repository.Token, data query.UpdateLicense) (repository.License, error) {
	entity, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	before, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, encodeErr
	}
	after, encodeErr := json.Marshal(data.Contract)
	if encodeErr != nil {
		return entity, encodeErr
	}
	{
		q := `INSERT INTO "license_audit" ("license_id", "contract", "what", "who", "with")
		      VALUES ($1, $2, $3, $4, $5)
		   RETURNING "when"`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, before,
			repository.Update, token.UserID, token.ID)
		if err := row.Scan(&entity.UpdatedAt); err != nil {
			return entity, err
		}
	}
	q := `UPDATE "license"
	         SET "contract" = $1, "updated_at" = $2
	       WHERE "id" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, after, entity.UpdatedAt, entity.ID)
	if err := row.Scan(&entity.UpdatedAt); err != nil {
		return entity, err
	}
	return entity, nil
}

// Delete TODO issue#docs
func (scope licenseManager) Delete(token *repository.Token, data query.DeleteLicense) (repository.License, error) {
	entity, readErr := scope.Read(token, query.ReadLicense{ID: data.ID})
	if readErr != nil {
		return entity, readErr
	}
	before, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, encodeErr
	}
	{
		q := `INSERT INTO "license_audit" ("license_id", "contract", "what", "who", "with")
		      VALUES ($1, $2, $3, $4, $5)
		   RETURNING "when"`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID, before,
			repository.Delete, token.UserID, token.ID)
		if err := row.Scan(&entity.DeletedAt); err != nil {
			return entity, err
		}
	}
	q := `UPDATE "license"
	         SET "updated_at" = $1, "deleted_at" = $2
	       WHERE "id" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.DeletedAt, entity.DeletedAt, entity.ID)
	if err := row.Scan(&entity.UpdatedAt); err != nil {
		return entity, err
	}
	return entity, nil
}
