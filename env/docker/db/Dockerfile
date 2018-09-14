FROM postgres:10-alpine AS db

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

ADD initdb.d /docker-entrypoint-initdb.d
