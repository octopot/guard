COMPOSE ?= docker-compose -f env/docker/compose/docker-compose.base.yml -f env/docker/compose/docker-compose.dev.yml -p guard


.PHONY: __env__
__env__:
	@(cp -nrf env/docker/compose/.env.example .env)
#|
#|                    ---
#|
.PHONY: config
config: __env__    #| Validate and view the Compose file.
	@($(COMPOSE) config)
#|
.PHONY: up
up: __env__        #| Builds, (re)creates, starts, and attaches to containers for services.
	@($(COMPOSE) up -d)
#|
.PHONY: fresh-up
fresh-up: __env__  #| Builds images before starting containers,
                   #| (re)creates containers even if their configuration and image haven't changed,
                   #| starts, and attaches to containers for a service.
	@($(COMPOSE) up -d --build --force-recreate)
#|
.PHONY: clean
clean: __env__     #| Removes stopped service containers.
	@($(COMPOSE) rm -f)
#|
.PHONY: down
down: __env__      #| Stops containers and removes them with networks.
	@($(COMPOSE) down)
#|
.PHONY: destroy
destroy: __env__   #| Stops containers and removes them with networks, volumes, and images created by `up`.
	@($(COMPOSE) down --volumes --rmi local)
#|
.PHONY: status
status: __env__    #| List containers and their status.
	@($(COMPOSE) ps)
#|
#|                    ---
#|
SERVICES = db \
           etcd \
           legacy \
           service \
           server

.PHONY: services
services:          #| Shows available services.
	@(echo 'available services:'; for container in $(SERVICES); do echo '-' $$container; done)

define service_tpl
#|
.PHONY: up-$(1)
up-$(1):           #| Builds, (re)creates, starts, and attaches to a container for the service $(1).
                   #| For example `make up-server`. See `make services`.
	@($$(COMPOSE) up -d $(1))
#|
.PHONY: container-$(1)
container-$(1):    #| Enter to a running container of the service $(1).
                   #| For example `make container-server`. See `make services`.
	@($$(COMPOSE) exec $(1) sh)
#|
.PHONY: start-$(1)
start-$(1):        #| Start an existing container of the service $(1).
                   #| For example `make start-server`. See `make services`.
	@($$(COMPOSE) start $(1))
#|
.PHONY: restart-$(1)
restart-$(1):      #| Restart a running container of the service $(1).
                   #| For example `make restart-server`. See `make services`.
	@($$(COMPOSE) restart $(1))
#|
.PHONY: stop-$(1)
stop-$(1):         #| Stop a running container of the service $(1) without removing them.
                   #| For example `make stop-server`. See `make services`.
	@($$(COMPOSE) stop $(1))
#|
.PHONY: log-$(1)
log-$(1):          #| View output from a container of the service $(1).
                   #| For example `make log-server`. See `make services`.
	@($$(COMPOSE) logs -f $(1))
#|
endef

render_service_tpl = $(eval $(call service_tpl,$(service)))
$(foreach service,$(SERVICES),$(render_service_tpl))
