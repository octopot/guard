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
		q := `INSERT INTO "license_audit" ("number", "contract", "what", "who", "with")
		      VALUES ($1, $2, $3, $4, $5)
		   RETURNING "when"`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.Number, before,
			repository.Create, token.UserID, token.ID)
		if err := row.Scan(&entity.CreatedAt); err != nil {
			return entity, err
		}
	}
	q := `INSERT INTO "license" ("number", "contract", "created_at")
	      VALUES ($1, $2, $3)
	   RETURNING "number", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, data.Number, after, entity.CreatedAt)
	if err := row.Scan(&entity.CreatedAt); err != nil {
		return entity, err
	}
	return entity, nil
}

// Read TODO issue#docs
func (scope licenseManager) Read(token *repository.Token, data query.ReadLicense) (repository.License, error) {
	entity, encoded := repository.License{Number: data.Number}, []byte(nil)
	q := `SELECT "contract", "created_at", "updated_at", "deleted_at"
	        FROM "license"
	       WHERE "number" = $1`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.Number)
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
	entity, readErr := scope.Read(token, query.ReadLicense{Number: data.Number})
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
		q := `INSERT INTO "license_audit" ("number", "contract", "what", "who", "with")
		      VALUES ($1, $2, $3, $4, $5)
		   RETURNING "when"`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.Number, before,
			repository.Update, token.UserID, token.ID)
		if err := row.Scan(&entity.UpdatedAt); err != nil {
			return entity, err
		}
	}
	q := `UPDATE "license"
	         SET "contract" = $1, "updated_at" = $2
	       WHERE "number" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, after, entity.UpdatedAt, entity.Number)
	if err := row.Scan(&entity.UpdatedAt); err != nil {
		return entity, err
	}
	return entity, nil
}

// Delete TODO issue#docs
func (scope licenseManager) Delete(token *repository.Token, data query.DeleteLicense) (repository.License, error) {
	entity, readErr := scope.Read(token, query.ReadLicense{Number: data.Number})
	if readErr != nil {
		return entity, readErr
	}
	before, encodeErr := json.Marshal(entity.Contract)
	if encodeErr != nil {
		return entity, encodeErr
	}
	{
		q := `INSERT INTO "license_audit" ("number", "contract", "what", "who", "with")
		      VALUES ($1, $2, $3, $4, $5)
		   RETURNING "when"`
		row := scope.conn.QueryRowContext(scope.ctx, q, entity.Number, before,
			repository.Delete, token.UserID, token.ID)
		if err := row.Scan(&entity.DeletedAt); err != nil {
			return entity, err
		}
	}
	q := `UPDATE "license"
	         SET "updated_at" = $1, "deleted_at" = $2
	       WHERE "number" = $3
	   RETURNING "updated_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.DeletedAt, entity.DeletedAt, entity.Number)
	if err := row.Scan(&entity.UpdatedAt); err != nil {
		return entity, err
	}
	return entity, nil
}
