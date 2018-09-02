
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
	             --force-rm \
	             .
