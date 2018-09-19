ifndef PACKAGE
$(error Please define PACKAGE variable)
endif

ifndef SECRET
$(error Please define SECRET variable)
endif

ifndef VERSION
$(error Please define VERSION variable)
endif


.PHONY: docker-build-db
docker-build-db:
	docker build \
	             -t guard-db:$(VERSION) \
	             -f env/docker/db/Dockerfile \
	             --force-rm \
	             env/docker/db/context

.PHONY: docker-build-etcd
docker-build-etcd:
	docker build \
	             -t guard-etcd:$(VERSION) \
	             -f env/docker/etcd/Dockerfile \
	             --force-rm \
	             env/docker/etcd/context

.PHONY: docker-build-legacy
docker-build-legacy:
	docker build \
	             -t guard-legacy:$(VERSION) \
	             -f env/docker/legacy/Dockerfile \
	             --force-rm \
	             env/docker/legacy/context

.PHONY: docker-build-server
docker-build-server:
	docker build \
	             -t guard-server:$(VERSION) \
	             -f env/docker/server/Dockerfile \
	             --force-rm \
	             env/docker/server/context

.PHONY: docker-build-service
docker-build-service:
	docker build \
	             -t guard-service:$(VERSION) \
	             -f env/docker/service/Dockerfile \
	             --build-arg PACKAGE=$(PACKAGE) \
	             --build-arg SECRET=$(SECRET) \
	             --force-rm \
	             .


.PHONY: docker-run-db
docker-run-db:
	docker run --rm -it \
	           -p 5432:5432 \
	           guard-db:$(VERSION)

.PHONY: docker-run-etcd
docker-run-etcd:
	docker run --rm -it \
	           -p 2379:2379 \
	           -p 2380:2380 \
	           guard-etcd:$(VERSION)

.PHONY: docker-run-service
docker-run-service:
	docker run --rm -it \
	           -p 8080:80 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           guard-service:$(VERSION) run --with-profiling --with-monitoring
