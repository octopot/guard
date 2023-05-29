---
code: null
id: MDU6SXNzdWUzODA3MzkzMDk=
databaseId: 380739309
number: 70
url: https://github.com/octopot/guard/issues/70
title: "problem with install command"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2018-11-14T15:03:15Z
updatedAt: 2018-11-24T16:49:37Z
lastEditedAt: null
closedAt: 2018-11-24T16:49:36Z
---

# problem with install command

Input:
```yml
account:
  name: QWEPADM
  users:
  - name: gorshkov
```

Expected: account + user + token

Obtained: account

Don't reproduced
```yml
account:
  name: QWEPADM
  users:
  - name: gorshkov
    tokens:
    - expired_at: ~
```
