# *************************************************************************** #
#                             Makefile
# *************************************************************************** #


NAME :=				bookWurms
COMPOSE_FILE :=		docker-compose.yaml
DOCKER_COMPOSE :=	docker compose -f $(COMPOSE_FILE)

DCKR_COMP_TEST :=	docker compose --profile test
GO_TEST :=			go test

all: go-mod up

re: down all

up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down

build:
	$(DOCKER_COMPOSE) build

build-new:
	$(DOCKER_COMPOSE) build --no-cache

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

.PHONY: all re up down build build-new logs ps stats volumes nuke-volumes

go-mod:
	cd Backend/ && go mod tidy

go-test:
	$(DCKR_COMP_TEST) up -d --wait db-test
	cd Backend/ && TEST_DATABASE_URL="host=localhost user=test_user password=test_password dbname=post_db_test port=5433 sslmode=disable" $(GO_TEST) ./...
	$(DCKR_COMP_TEST) down

go-test-verbose: GO_TEST := go test -v -count=1
go-test-verbose: go-test

help:
	@echo "Options:"
	@echo "  all (default)   - alias for: up + go-mod"
	@echo "  re              - alias for: down + all"
	@echo "  up              - Build and start the containers"
	@echo "  down            - Stop and remove the containers"
	@echo "  build           - Build the containers"
	@echo "  logs            - Show logs of the containers"
	@echo "  ps              - List running containers"
	@echo "  stats           - Show resource usage statistics of the containers"
	@echo "  volumes         - List Docker volumes"
	@echo "  nuke-volumes    - DELETE ALL the volumes with down -v"
	@echo "  go-mod          - Run go mod tidy in the Backend"
	@echo "  go-test         - Run go test ./... in the Backend"
	@echo "  go-test-verbose - Run go test -v -count=1 ./... in the Backend"
	@echo "  help            - Show this help"

.PHONY: go-mod go-test go-test-verbose help