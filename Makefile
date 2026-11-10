# *************************************************************************** #
#                             Makefile
# *************************************************************************** #


NAME :=				bookWurms
COMPOSE_FILE :=		docker-compose.yml
DOCKER_COMPOSE :=	docker compose -f $(COMPOSE_FILE)

all: go-mod up

re: down all

up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down

build:
	$(DOCKER_COMPOSE) build

logs:
	$(DOCKER_COMPOSE) logs

ps:
	$(DOCKER_COMPOSE) ps

stats:
	$(DOCKER_COMPOSE) stats


volumes:
	$(DOCKER_COMPOSE) volumes

nuke-volumes:
	@echo "Delete all volumes(db, caddy)? [y/N] " && read ans && [ "$$ans" = y ]
	$(DOCKER_COMPOSE) down -v

.PHONY: all re up down build logs ps stats volumes nuke-volumes

go-mod:
	cd Backend && go mod tidy

help:
	@echo "Options:"
	@echo "  all (default) - alias for: up + go-mod"
	@echo "  re            - alias for: down + all"
	@echo "  up            - Build and start the containers"
	@echo "  down          - Stop and remove the containers"
	@echo "  build         - Build the containers"
	@echo "  logs          - Show logs of the containers"
	@echo "  ps            - List running containers"
	@echo "  stats         - Show resource usage statistics of the containers"
	@echo "  volumes       - List Docker volumes"
	@echo "  nuke-volumes  - DELETE ALL the volumes with down -v"
	@echo "  go-mod        - Run go mod tidy in the backend"
	@echo "  help          - Show this help"

.PHONY: go-mod help