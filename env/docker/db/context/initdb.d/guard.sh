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
EOSQL
