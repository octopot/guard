COMPOSE ?= docker-compose -f env/docker/compose/docker-compose.base.yml -f env/docker/compose/docker-compose.dev.yml -p guard


.PHONY: env
env:
	cp -nrf env/docker/compose/.env.example .env


.PHONY: config
config: env
	$(COMPOSE) config

.PHONY: up
up: env
	$(COMPOSE) up -d

.PHONY: fresh-up
fresh-up: env
	$(COMPOSE) up -d --build --force-recreate

.PHONY: clear
clear: env
	$(COMPOSE) rm -f

.PHONY: down
down: env
	$(COMPOSE) down

.PHONY: destroy
destroy: env
	$(COMPOSE) down --volumes --rmi local

.PHONY: status
status: env
	$(COMPOSE) ps
