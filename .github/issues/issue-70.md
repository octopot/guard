---
id: 70
database_id: 380739309
node_id: MDU6SXNzdWUzODA3MzkzMDk=
status: closed
title: "problem with install command"
labels: ["type:bug"]
url: https://github.com/octopot/guard/issues/70
created_at: 2018-11-14T15:03:15Z
updated_at: 2018-11-24T16:49:37Z
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
