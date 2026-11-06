# *************************************************************************** #
#                             Makefile
# *************************************************************************** #


NAME :=				ft_transcendence
COMPOSE_FILE :=		docker-compose.yml
DOCKER_COMPOSE :=	docker compose -f $(COMPOSE_FILE)

all: go-mod up

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

re: all

go-mod:
	cd Backend/src && go mod tidy

# Local install so the editor finds the packages/types (also updates package-lock.json)
npm-install:
	cd Frontend && npm install

# Type-check the frontend inside the running container
typecheck:
	$(DOCKER_COMPOSE) exec frontend npm run typecheck

caddy-reload:
	$(DOCKER_COMPOSE) exec -w /etc/caddy caddy caddy reload

help:
	@echo "Options:"
	@echo "  all (default) - alias for up + go-mod"
	@echo "  up            - Build and start the containers"
	@echo "  down          - Stop and remove the containers"
	@echo "  build         - Build the containers"
	@echo "  logs          - Show logs of the containers"
	@echo "  ps            - List running containers"
	@echo "  stats         - Show resource usage statistics of the containers"
	@echo "  volumes       - List Docker volumes"
	@echo "  nuke-db       - DELETE the database volume with down -v"
	@echo "  re            - alias for all"
	@echo "  go-mod        - Run go mod tidy in the backend"
	@echo "  npm-install   - Install frontend packages locally (for the editor)"
	@echo "  typecheck     - Type-check the frontend in its container"
	@echo "  caddy-reload  - Reload the Caddy configuration"
	@echo "  help          - Show this help"

.PHONY: all up down build logs ps stats volumes nuke-db re go-mod npm-install typecheck caddy-reload help