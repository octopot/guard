FROM php:latest AS build

ADD . /opt/legacy/


FROM php:fpm-alpine AS app

LABEL maintainer="Kamil Samigullin <kamil@samigullin.info>"

COPY --from=build /opt/legacy /var/www/html
