ifndef PACKAGE
$(error Please define PACKAGE variable)
endif

ifndef VERSION
$(error Please define VERSION variable)
endif


.PHONY: docker-build-app
docker-build-app:
	docker build \
	             -t paymaster-app:$(VERSION) \
	             -f env/docker/app/Dockerfile \
	             --force-rm \
	             env/docker/app/context

.PHONY: docker-build-server
docker-build-server:
	docker build \
	             -t paymaster-server:$(VERSION) \
	             -f env/docker/server/Dockerfile \
	             --force-rm \
	             env/docker/server/context

.PHONY: docker-build-service
docker-build-service:
	docker build \
	             -t paymaster-service:$(VERSION) \
	             -f env/docker/service/Dockerfile \
	             --build-arg PACKAGE=$(PACKAGE) \
	             --force-rm \
	             .


.PHONY: docker-run-service
docker-run-service:
	docker run --rm -it \
	           -p 8080:80 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           paymaster-service:$(VERSION) run --with-profiling --with-monitoring
