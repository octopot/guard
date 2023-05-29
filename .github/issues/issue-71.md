---
code: null
id: MDU6SXNzdWUzODIyMTIxMjQ=
databaseId: 382212124
number: 71
url: https://github.com/octopot/guard/issues/71
title: "review init script and migrations"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-11-19T13:21:50Z
updatedAt: 2018-11-23T09:37:39Z
lastEditedAt: null
closedAt: 2018-11-23T09:37:39Z
---

# review init script and migrations

- https://github.com/kamilsk/guard/blob/8d8aaabc0498264103d1d6dbed03bcb21cd85284/pkg/storage/migrations/postgres/1_prepare.sql#L3
- move `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` to https://github.com/kamilsk/guard/blob/master/env/docker/db/context/initdb.d/guard.sh
