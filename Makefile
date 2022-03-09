.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: compose-up
compose-up: compose-down compose-build docker-scan  ## Create and start containers
	docker-compose up

.PHONY: compose-down
compose-down: ## Stop and remove containers, networks, images, and volumes
	docker-compose down --remove-orphans

.PHONY: compose-restart
compose-restart: compose-up ## restart services

.PHONY: compose-tail
compose-tail: ## Tail output from containers
	docker-compose logs -f

.PHONY: compose-build
compose-build: ## Build or rebuild services
	docker-compose build --no-cache

.PHONY: compose-top
compose-top: ## Display the running processes
	docker-compose top

.PHONY: compose-ps
compose-ps: ## List containers
	docker-compose ps

.PHONY: docker-scan
compose-ps: ## Scan docker image after building
	docker scan api




