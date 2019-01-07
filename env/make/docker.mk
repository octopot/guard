ifndef VERSION
$(error Please define VERSION variable)
endif


.PHONY: docker-build
docker-build:
	docker build -f env/docker/service/Dockerfile \
	             -t kamilsk/guard:$(VERSION) \
	             -t kamilsk/guard:latest \
	             -t quay.io/kamilsk/guard:$(VERSION) \
	             -t quay.io/kamilsk/guard:latest \
	             --force-rm --no-cache --pull --rm \
	             .

.PHONY: docker-push
docker-push:
	docker push kamilsk/guard:$(VERSION)
	docker push kamilsk/guard:latest
	docker push quay.io/kamilsk/guard:$(VERSION)
	docker push quay.io/kamilsk/guard:latest

.PHONY: docker-refresh
docker-refresh:
	docker images --all \
	| grep '^kamilsk\/guard\s\+' \
	| awk '{print $$3}' \
	| xargs docker rmi -f &>/dev/null || true
	docker pull kamilsk/guard:$(IMAGE_VERSION)



.PHONY: publish
publish: docker-build docker-push



.PHONY: docker-start
docker-start:
	docker run --rm -it \
	           --env-file env/.env.example \
	           --name guard-dev \
	           -p 8080:8080 \
	           -p 8090:8090 \
	           -p 8091:8091 \
	           -p 8092:8092 \
	           -p 8093:8093 \
	           guard-service:$(VERSION) run --with-profiling --with-monitoring --with-grpc-gateway
