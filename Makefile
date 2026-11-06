# *************************************************************************** #
#                             Makefile
# *************************************************************************** #


NAME :=				ft_transcendence
COMPOSE_FILE :=		docker-compose.yml
DOCKER_COMPOSE :=	docker compose -f $(COMPOSE_FILE)

all: up

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

nuke-db:
	@echo "Delete the database volume? [y/N] " && read ans && [ "$$ans" = y ]
	$(DOCKER_COMPOSE) down -v

re: up


help:
	@echo "Options:"
	@echo "  all (default) - alias for up"
	@echo "  up            - Build and start the containers"
	@echo "  down          - Stop and remove the containers"
	@echo "  build         - Build the containers"
	@echo "  logs          - Show logs of the containers"
	@echo "  ps            - List running containers"
	@echo "  stats         - Show resource usage statistics of the containers"
	@echo "  volumes       - List Docker volumes"
	@echo "  nuke-db       - DELETE the database volume with down -v"
	@echo "  re            - alias for up"
	@echo "  help          - Show this help"

.PHONY: all up down build logs ps stats volumes nuke-db re help