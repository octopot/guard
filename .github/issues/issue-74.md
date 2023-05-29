---
code: null
id: MDU6SXNzdWUzODQwNDkwNjQ=
databaseId: 384049064
number: 74
url: https://github.com/octopot/guard/issues/74
title: "bug with unexpected domain rate unit"
issueType: null
labels:
  - "type:critical"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-11-25T05:23:59Z
updatedAt: 2018-11-25T05:29:55Z
lastEditedAt: null
closedAt: 2018-11-25T05:29:55Z
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
