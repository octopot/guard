_base   = docker-compose -p guard -f env/docker/compose/docker-compose.base.yml
COMPOSE = $(_base) -f env/docker/compose/docker-compose.dev.yml


.PHONY: __env__
__env__:
	@(cp -nrf env/docker/compose/.env.example .env)
#|
#|                    --- Docker Compose' generic commands
#|
.PHONY: ci
ci:                #| Switch docker compose to CI/CD configuration.
	$(eval COMPOSE = $(_base) -f env/docker/compose/docker-compose.ci.yml)
	@(echo $(COMPOSE))
#|
.PHONY: demo
demo:              #| Switch docker compose to demo configuration.
	$(eval COMPOSE = $(_base) -f env/docker/compose/docker-compose.demo.yml)
	@(echo $(COMPOSE))
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
#|                    --- Service-specific commands
#|
SERVICES = db \
           legacy \
           migration \
           service \
           server \
           spec

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
.PHONY: fresh-up-$(1)
fresh-up-$(1):     #| Builds images before starting containers,
                   #| (re)creates containers even if their configuration and image haven't changed,
                   #| starts, and attaches to containers for the service $(1).
	@($$(COMPOSE) up -d --build --force-recreate $(1))
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
#|                    --- Database-specific commands
#|
.PHONY: psql
psql: __env__      #| Connect to the database.
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c psql')
#|
.PHONY: backup
backup: __env__    #| Backup the database.
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "pg_dump --format=custom --file=/tmp/db.dump $${POSTGRES_DB}"')
	@(docker cp $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/db.dump ./env/docker/db/)
	@($(COMPOSE) exec db rm /tmp/db.dump)
	@(ls -l ./env/docker/db/db.dump)
#|
.PHONY: restore
restore: __env__   #| Restore the database.
	@(docker cp ./env/docker/db/reset.sql $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@(docker cp ./env/docker/db/db.dump $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "psql $${POSTGRES_DB} < /tmp/reset.sql"')
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "pg_restore -Fc -d $${POSTGRES_DB} /tmp/db.dump"')
	@($(COMPOSE) exec db rm /tmp/reset.sql /tmp/db.dump)
#|
.PHONY: truncate
truncate: __env__  #| Truncate the database tables.
	@(docker cp ./env/docker/db/truncate.sql $$(make status | tail +3 | awk '{print $$1}' | grep _db_ | head -1):/tmp/)
	@($(COMPOSE) exec db /bin/sh -c 'su - postgres -c "psql $${POSTGRES_DB} < /tmp/truncate.sql"')
	@($(COMPOSE) exec db rm /tmp/truncate.sql)
