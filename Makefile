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

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

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

bazel-test: ### runs the test cases with bazel
	@echo "::bazel-test"; \
    bazelisk test \
    	--action_env=DB_TYPE=$(DB_TYPE) \
    	--action_env=DB_USERNAME=$(DB_USERNAME) \
    	--action_env=DB_PASSWORD=$(DB_PASSWORD) \
    	--action_env=DB_HOSTNAME=$(DB_HOSTNAME) \
    	--action_env=DB_PORT=$(DB_PORT) \
    	--action_env=POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
    	--action_env=SERVER_HOST=$(SERVER_HOST) \
    	--action_env=GRPC_SERVER_PORT=$(GRPC_SERVER_PORT) \
    	--action_env=GRPC_GATEWAY_SERVER_PORT=$(GRPC_GATEWAY_SERVER_PORT) \
        --platform_suffix="bazel-test" \
        --@io_bazel_rules_go//go/config:race \
        --define cluster=$CLUSTER \
        --define namespace=$NAMESPACE \
        --test_tag_filters=fast \
        --build_tag_filters=fast \
        --test_output=errors \
        --nocache_test_results \
        //...
.PHONY: bazel-test
