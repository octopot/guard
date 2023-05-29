---
code: null
id: MDU6SXNzdWUzODAwNzIzODQ=
databaseId: 380072384
number: 69
url: https://github.com/octopot/guard/issues/69
title: "bug with grpc client"
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
createdAt: 2018-11-13T06:02:17Z
updatedAt: 2018-11-13T20:15:00Z
lastEditedAt: 2018-11-13T06:03:33Z
closedAt: 2018-11-13T20:14:59Z
---

# bug with grpc client

```sh
$ guardctl install -f /opt/guard/install.yml -t 1m
rpc error: code = Internal desc = something happen: trying to get connection
$ netstat -nl | grep :80
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN
tcp        0      0 0.0.0.0:8080            0.0.0.0:*               LISTEN
tcp        0      0 127.0.0.1:8090          0.0.0.0:*               LISTEN
tcp        0      0 127.0.0.1:8092          0.0.0.0:*               LISTEN
tcp        0      0 127.0.0.1:8093          0.0.0.0:*               LISTEN
```

<!-- 2018-11-13T09:18Z https://github.com/octopot/guard/issues/69#issuecomment-438193700
telnet is ok, problem can be reproduced at Ubuntu
-->

<!-- 2018-11-13T09:22Z https://github.com/octopot/guard/issues/69#issuecomment-438194740
problem is actual for grpc gateway too

```
{"error":"something happen: trying to get connection","message":"something happen: trying to get connection","code":13}
```
-->

<!-- 2018-11-13T20:14Z https://github.com/octopot/guard/issues/69#issuecomment-438420444
problems related to database, false positive
-->
