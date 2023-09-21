# DOCKER TASKS =========================================================================================================
# This is the default. It can be overridden in the main Makefile after
# including docker.mk

up: ## Start the docker-compose environment
	@docker compose -f docker-compose.yaml \
		up -d --remove-orphans --build

down: confirm ## Down docker compose
	@docker compose -f docker-compose.yaml \
		down --remove-orphans
