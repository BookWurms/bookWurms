# *************************************************************************** #
#                             Makefile
# *************************************************************************** #


NAME :=				bookWurms
COMPOSE_FILE :=		docker-compose.yaml
DOCKER_PROD :=		docker compose -f $(COMPOSE_FILE)
DOCKER_ALL :=		$(DOCKER_PROD) --profile "*"
DOCKER_DEV :=		$(DOCKER_PROD) --profile dev
DOCKER_TEST :=		$(DOCKER_PROD) --profile test

GO_TEST :=			go test


# ---------------------------------------------------------------------------
.PHONY: all re 
# ---------------------------------------------------------------------------

# TODO: change to docker-up-prod at the end
all: go-tidy docker-up-dev

re: docker-down all


# ---------------------------------------------------------------------------
.PHONY: docker-up-dev docker-up-prod docker-down docker-build-dev-no-cache 
.PHONY: docker-logs docker-ps docker-nuke-volumes
# ---------------------------------------------------------------------------

# This is the final project result 
docker-up-prod:
	$(DOCKER_PROD) up --build -d

docker-up-dev:
	$(DOCKER_DEV) up --build -d

docker-down:
	$(DOCKER_ALL) down --remove-orphans

docker-logs:
	$(DOCKER_ALL) logs

docker-ps:
	$(DOCKER_ALL) ps

docker-nuke-volumes:
	@echo "Delete all volumes(postgres-db, caddy)? [y/N] " && read ans && [ "$$ans" = y ]
	$(DOCKER_ALL) down -v

docker-build-dev-no-cache:
	$(DOCKER_DEV) build --no-cache 


# ---------------------------------------------------------------------------
.PHONY: go-tidy go-test go-test-verbose
# ---------------------------------------------------------------------------

# update go.mod and go.sum
go-tidy:
	cd Backend/ && go mod tidy

go-test:
	$(DOCKER_TEST) up -d --wait postgres-db-test
	cd Backend/ && TEST_DATABASE_URL="host=localhost user=test_user password=test_password dbname=post_db_test port=5433 sslmode=disable" $(GO_TEST) ./...

go-test-verbose: GO_TEST := go test -v -count=1
go-test-verbose: go-test


# ---------------------------------------------------------------------------
.PHONY: help
# ---------------------------------------------------------------------------

help:
	@echo "Options:"
	@echo "  all (default)       - alias:  go-tidy + docker-up-dev"
	@echo "  re                  - alias:  docker-down + all"
	@echo "  docker-up-prod      - run:   in production mode"
	@echo "  docker-up-dev       - run:   prod + dev"
	@echo "  docker-down         - stop+remove: prod + dev + test"
	@echo "  docker-logs         - logs:  prod + dev + test"
	@echo "  docker-ps           - list:  prod + dev + test"
	@echo "  docker-nuke-volumes - remove: ALL volumes"
	@echo "  docker-build-dev-no-cache"
	@echo "  go-tidy             - run:   go mod tidy"
	@echo "  go-test             - run:   go test ./..."
	@echo "  go-test-verbose     - run:   go test -v -count=1 ./..."
	@echo "  help                - print: this help"
