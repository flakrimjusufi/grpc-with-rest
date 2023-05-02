include .env
export

CLUSTER=$1
NAMESPACE=$2

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

docker-up: ### Run docker-compose
	docker-compose up --build -d db app && docker-compose logs -f
.PHONY: docker-up

docker-restart: docker-down docker-up ### Restart containers
.PHONY: docker-restart

docker-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: docker-down

docker-clean-volumes: ### clean volumes created by docker
	echo "::docker-clean-volumes";
	echo ">>> Deleting unused volumes";
	docker system prune --volumes -f;
.PHONY: docker-clean-volumes

docker-rm-volume: ### remove docker volume
	docker volume rm neocharge_database_data
.PHONY: docker-rm-volume

docker-clean-images: ## deletes untagged images
	echo "docker-clean-images";
	echo ">>> Deleting untagged images";
	docker rmi `docker images -f dangling=true -q`
.PHONY: docker-clean-images

### BAZEL COMMANDS
bazel-clean: ### clean bazel cached files
	@echo "::bazel-clean";
	@bazelisk clean;
.PHONY: bazel-clean

bazel-setup: ### creates the setup for bazel
	@echo "::bazel-setup";
	@bazelisk run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:gazelle
.PHONY: bazel-setup

bazel-update-deps: ### updates the dependencies if a dependency has changed/added
	@echo "::bazel-update-deps";
	@bazelisk run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:gazelle -- update;
	@bazelisk run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:gazelle -- update-repos -from_file=go.mod;
.PHONY: bazel-update-deps

bazel-run: ### runs the project with bazel
	@echo "::bazel-run";
	@DB_USERNAME=$(DB_USERNAME) \
    DB_PASSWORD=$(DB_PASSWORD) \
    DB_DATABASE=$(DB_DATABASE) \
    DB_HOSTNAME=$(DB_HOSTNAME) \
    DB_PORT=$(DB_PORT) \
    DB_TYPE=$(DB_TYPE) \
    POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
    SERVER_HOST=$(SERVER_HOST) \
    GRPC_SERVER_PORT=$(GRPC_SERVER_PORT) \
    GRPC_GATEWAY_SERVER_PORT=$(GRPC_GATEWAY_SERVER_PORT) \
    bazelisk run //:gazelle
.PHONY: bazel-run
