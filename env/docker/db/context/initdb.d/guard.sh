#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" <<-EOSQL
    \c "${POSTGRES_DB}";
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE TABLE "license_token" (
      "license"    UUID      NOT NULL,
      "token"      UUID      NOT NULL,
      "created_at" TIMESTAMP NOT NULL             DEFAULT now(),
      UNIQUE ("license", "token")
    );
    INSERT INTO "license_token" ("license", "token")
    VALUES
      ('10000000-2000-4000-8000-160000000001', '10000000-2000-4000-8000-160000000002'),
      ('10000000-2000-4000-8000-160000000001', '10000000-2000-4000-8000-160000000003'),
      ('10000000-2000-4000-8000-160000000004', '10000000-2000-4000-8000-160000000005');
EOSQL
