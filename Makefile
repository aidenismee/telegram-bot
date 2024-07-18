.DEFAULT_GOAL := help

GO111MODULE := on
HOST ?= localhost:8080

# Generates a help message. Borrowed from https://github.com/pydanny/cookiecutter-djangopackage.
help: ## Display this help message
	@echo "Please use \`make <target>' where <target> is one of"
	@perl -nle'print $& if m{^[\.a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

depends: ## Install & build dependencies
	go get ./...
	go build -o /dev/null ./...
	go mod tidy

dev.provision:  ## Provision dev environment
	docker-compose up -d
	#scripts/waitdb.sh
	@$(MAKE) migrate

migrate: ## Run database migrations
	go run cmd/migration/main.go

run: ## Run the server
	go run cmd/api/main.go

clean: ## Clean up
	rm -rf ./main ./*.out