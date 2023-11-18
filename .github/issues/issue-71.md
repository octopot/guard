---
id: 71
database_id: 382212124
node_id: MDU6SXNzdWUzODIyMTIxMjQ=
status: closed
title: "review init script and migrations"
labels: ["type:refactoring"]
url: https://github.com/octopot/guard/issues/71
created_at: 2018-11-19T13:21:50Z
updated_at: 2018-11-23T09:37:39Z
---

# review init script and migrations

- https://github.com/kamilsk/guard/blob/8d8aaabc0498264103d1d6dbed03bcb21cd85284/pkg/storage/migrations/postgres/1_prepare.sql#L3
- move `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` to https://github.com/kamilsk/guard/blob/master/env/docker/db/context/initdb.d/guard.sh
