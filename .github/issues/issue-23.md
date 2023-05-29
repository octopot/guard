---
code: null
id: MDU6SXNzdWUzNjA1NzA1NDI=
databaseId: 360570542
number: 23
url: https://github.com/octopot/guard/issues/23
title: "etcd storage layer"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2018-09-15T19:51:31Z
updatedAt: 2018-10-27T19:48:46Z
lastEditedAt: 2018-09-26T11:44:48Z
closedAt: null
---

# etcd storage layer

- [ ] executor
- [ ] query
- [ ] types

<!-- 2018-10-02T13:06Z https://github.com/octopot/guard/issues/23#issuecomment-426265977
```yaml

  etcd:
    image: guard-etcd:latest
    env_file: ../etcd/.env

  etcd:
    build:
      context:    ../etcd/context
      dockerfile: ../Dockerfile
    ports:
    - ${LOCAL_ETCD_PORT}:2379

```
-->
