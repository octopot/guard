IMAGE_VERSION := latest
PACKAGE       := github.com/kamilsk/guard


.PHONY: docker-build-app
docker-build-app:
	docker build \
	             -t paymaster-app \
	             -f env/docker/app/Dockerfile \
	             --force-rm \
	             env/docker/app/context

.PHONY: docker-build-server
docker-build-server:
	docker build \
	             -t paymaster-server \
	             -f env/docker/server/Dockerfile \
	             --force-rm \
	             env/docker/server/context

.PHONY: docker-build-service
docker-build-service:
	docker build \
	             -t paymaster-service \
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
	           paymaster-service run --with-profiling --with-monitoring
