#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" <<-EOSQL
    \c "${POSTGRES_DB}";
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE FUNCTION update_timestamp()
      RETURNS TRIGGER AS \$update_timestamp\$
    BEGIN
      IF NEW.* IS DISTINCT FROM OLD.*
      THEN
        NEW.updated_at := current_timestamp;
        RETURN NEW;
      ELSE
        RETURN OLD;
      END IF;
    END;
    \$update_timestamp\$
    LANGUAGE plpgsql;
    CREATE FUNCTION ignore_update()
      RETURNS TRIGGER AS \$ignore_update\$
    BEGIN
      RETURN OLD;
    END;
    \$ignore_update\$
    LANGUAGE plpgsql;
    CREATE TABLE "token" (
      "id"         UUID      NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
      "user_id"    UUID      NOT NULL,
      "expired_at" TIMESTAMP NULL                 DEFAULT NULL,
      "created_at" TIMESTAMP NOT NULL             DEFAULT now()
    );
    CREATE TRIGGER "immutable_token"
      BEFORE UPDATE
      ON "token"
      FOR EACH ROW EXECUTE PROCEDURE ignore_update();
    INSERT INTO "token" ("id", "user_id") VALUES ('10000000-2000-4000-8000-160000000000', '10000000-2000-4000-8000-160000000001');
    CREATE TABLE "license_user" (
      "license"    UUID      NOT NULL,
      "user_id"    UUID      NOT NULL,
      "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
      UNIQUE ("license", "user_id")
    );
    INSERT INTO "license_user" ("user_id", "license")
    VALUES
      ('10000000-2000-4000-8000-160000000001', '10000000-2000-4000-8000-160000000002'),
      ('10000000-2000-4000-8000-160000000003', '10000000-2000-4000-8000-160000000002'),
      ('10000000-2000-4000-8000-160000000005', '10000000-2000-4000-8000-160000000004');
EOSQL
