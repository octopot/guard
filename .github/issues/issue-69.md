---
id: 69
database_id: 380072384
node_id: MDU6SXNzdWUzODAwNzIzODQ=
status: closed
title: "bug with grpc client"
labels: ["type:critical"]
url: https://github.com/octopot/guard/issues/69
created_at: 2018-11-13T06:02:17Z
updated_at: 2018-11-13T20:15:00Z
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
