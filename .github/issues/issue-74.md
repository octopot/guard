---
id: 74
database_id: 384049064
node_id: MDU6SXNzdWUzODQwNDkwNjQ=
status: closed
title: "bug with unexpected domain rate unit"
labels: ["type:critical"]
url: https://github.com/octopot/guard/issues/74
created_at: 2018-11-25T05:23:59Z
updated_at: 2018-11-25T05:29:55Z
---

# bug with unexpected domain rate unit

1. register license without rate section
2. try to read the license

expected: read succeed
obtained:
```json
{
  "error": "unexpected domain rate unit ",
  "message": "unexpected domain rate unit ",
  "code": 13
}
```
