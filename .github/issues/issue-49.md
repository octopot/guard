---
id: 49
database_id: 365260079
node_id: MDU6SXNzdWUzNjUyNjAwNzk=
status: closed
title: "delete operation of license server is not idempotent"
labels: ["type:critical"]
url: https://github.com/octopot/guard/issues/49
created_at: 2018-09-30T20:48:48Z
updated_at: 2018-10-01T08:56:37Z
---

# delete operation of license server is not idempotent

- [x] if license already deleted return that fact
- [x] otherwise, mark it as deleted
