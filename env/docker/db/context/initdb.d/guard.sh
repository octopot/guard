#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" <<-EOSQL
    \c "${POSTGRES_DB}";
EOSQL
