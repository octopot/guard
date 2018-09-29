-- +migrate Up

CREATE TABLE "license" (
  "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "account_id" UUID      NOT NULL,
  "contract"   JSONB     NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  "deleted_at" TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "license_audit" (
  "id"         BIGSERIAL PRIMARY KEY,
  "license_id" UUID      NOT NULL,
  "contract"   JSONB     NOT NULL,
  "what"       ACTION    NOT NULL,
  "who"        UUID      NOT NULL,
  "when"       TIMESTAMP NOT NULL,
  "with"       UUID      NOT NULL
);

CREATE TRIGGER "license_updated"
  BEFORE UPDATE
  ON "license"
  FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER "immutable_license_audit"
  BEFORE UPDATE
  ON "license_audit"
  FOR EACH ROW EXECUTE PROCEDURE ignore_update();



-- +migrate Down

DROP TRIGGER "immutable_license_audit" ON "license_audit";

DROP TRIGGER "license_updated" ON "license";

DROP TABLE "license_audit";

DROP TABLE "license";
