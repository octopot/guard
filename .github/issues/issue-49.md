---
code: null
id: MDU6SXNzdWUzNjUyNjAwNzk=
databaseId: 365260079
number: 49
url: https://github.com/octopot/guard/issues/49
title: "delete operation of license server is not idempotent"
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
createdAt: 2018-09-30T20:48:48Z
updatedAt: 2018-10-01T08:56:37Z
lastEditedAt: 2018-10-01T08:56:21Z
closedAt: 2018-10-01T08:56:37Z
---

# delete operation of license server is not idempotent

- [x] if license already deleted return that fact
- [x] otherwise, mark it as deleted
