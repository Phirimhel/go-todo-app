include .env
export

export PROJECT_ROOT=$(shell pwd)

# Color Scheme
RED    := \033[0;31m
YELLOW := \033[0;33m
GREEN  := \033[0;32m
NC     := \033[0m 


# run app
app-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go

# deploy 
app-deploy:
	@docker compose up -d --build todoapp


app-undeploy:
	@docker compose down todoapp

# logs 
clear-logs:
	@echo "Clean all logs? ${YELLOW}Warning:${NC} Risk of losing all log files. [y/N]: \c"; \
	read ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf out/logs/* && \
		echo "${GREEN}All log files were wiped${NC}"; \
	else \
		echo "Logs cleaning aborted"; \
	fi
# postgres 
 
env-up: 
	docker compose up -d  todoapp-postgres

env-down: 
	docker compose down  todoapp-postgres

env-cleanup:
	@read -p "Clean all volume files? ${YELLOW}Warning:${NC} Risk of data loss. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down && \
		rm -rf out/pgdata && \
		echo "${GREEN}Files were wiped${NC}"; \
	else \
		echo "Cleanup aborted"; \
	fi


#migrations

migrate-create: 
	@if [ -z "$(seq)" ]; then \
		echo "${RED}Error:${NC} Migration name is missing."; \
		echo "${YELLOW}Usage:${NC} make migrate-create seq=<migration_name>"; \
		echo "Example: make migrate-create seq=create_users_table"; \
		exit 1; \
	fi	
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"


migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-force:
	@docker compose run --rm todoapp-postgres-migrate \
		-path migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		force 1

migrate-action:
	@if docker compose run --rm todoapp-postgres-migrate \
		-path migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action); then \
			echo "$(GREEN)Migration $(action) completed successfully$(NC)"; \
	else \
			echo "$(RED)Migration $(action) failed$(NC)"; \
			exit 1; \
	fi

migrate-version:
	@docker compose run --rm todoapp-postgres-migrate \
		-path migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		version 


# ports
env-port-foward:
	@docker compose up -d todoapp-port-forwarder 
	
env-port-close:
	@docker compose down todoapp-port-forwarder 


# docer compose ps 
ps:
	@docker compose ps


# swagger
swagger-gen: 
	@docker compose run --rm swagger \
		init \
		-g cmd/todoapp/main.go \
		-o docs \
		--parseInternal \
		--parseDependency
