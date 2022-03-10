.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: compose-up
compose-up: compose-down compose-build ## Create and start containers
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
docker-scan: ## Scan docker image after building
	docker scan api

.PHONY: setup-infra
setup-infra: ## Setting up infra
	@echo 'Setting up infrausing helmcharts'
	helm repo add bitnami https://charts.bitnami.com/bitnami

.PHONY: install-mongodb
install-mongodb: setup-infra ## Install  mongodb
	@echo 'Installing mongodb using helmcharts'
	helm upgrade --install mongodb bitnami/mongodb -f mongodb-k8-manifest/values.yaml
	sleep 20s

.PHONY: delete-mongodb
delete-mongodb: ## Delete  mongodb
	@echo 'Deleting mongodb using helmcharts'
	helm delete mongodb

.PHONY: install-nginx-controller
install-nginx-controller: setup-infra ## Install nginx controller
	@echo 'Installing nginx controller using helmcharts'
	helm upgrade --install nginx-ingress-controller bitnami/nginx-ingress-controller
	sleep 20s

.PHONY: install-nginx-controller
delete-nginx-controller: ## Delete nginx controller
	@echo 'Deleting nginx controller using helmcharts'
	helm delete nginx-ingress-controller

.PHONY: install-api
install-api: install-nginx-controller install-mongodb ## install api
	@echo 'Installing the rest-api'
	helm upgrade --install api ./api-k8-helm-manifest -f api-k8-helm-manifest/values.yaml

.PHONY: uninstall-api
uninstall-api: delete-nginx-controller delete-app ## uninstall api
	@echo 'Installing the rest-api'
	helm upgrade --install api ./api-k8-helm-manifest -f api-k8-helm-manifest/values.yaml

delete-app: delete-mongodb ## upgrade rest-api
	@echo 'Deleting the api'
	helm delete api