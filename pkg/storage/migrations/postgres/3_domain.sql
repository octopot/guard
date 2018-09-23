-- +migrate Up

CREATE TABLE "license" (
  "number"       UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  "active_since" TIMESTAMP NULL                 DEFAULT NULL,
  "active_until" TIMESTAMP NULL                 DEFAULT NULL,
  -- {"rate_limits":"%d (rps|rpm|rph|rpw)", "request_limits": uint, "workplace_limits": uint}
  "contract"     JSONB     NULL                 DEFAULT NULL,
  "created_at"   TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at"   TIMESTAMP NULL                 DEFAULT NULL,
  "deleted_at"   TIMESTAMP NULL                 DEFAULT NULL
);

CREATE TABLE "license_audit" (
  "id"           BIGSERIAL PRIMARY KEY,
  "number"       UUID      NOT NULL,
  "what"         ACTION    NOT NULL,
  "when"         TIMESTAMP NOT NULL,
  "who"          UUID      NOT NULL,
  -- before
  "active_since" TIMESTAMP NULL,
  "active_until" TIMESTAMP NULL,
  "contract"     JSONB     NULL
);

CREATE TABLE "license_user" (
  "license"    UUID      NOT NULL,
  "user_id"    UUID      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  UNIQUE ("license", "user_id")
);

CREATE TABLE "license_workplace" (
  "license"    UUID      NOT NULL,
  "workplace"  UUID      NOT NULL,
  "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
  "updated_at" TIMESTAMP NULL                 DEFAULT NULL,
  UNIQUE ("license", "workplace")
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

DROP TRIGGER "immutable_license_audit"
ON "license_audit";

DROP TRIGGER "license_updated"
ON "license";

DROP TABLE "license_workplace";

DROP TABLE "license_user";

DROP TABLE "license_audit";

DROP TABLE "license";
