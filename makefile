APP_NAME = go-hello-server
PKG     = github.com/darahayes/$(APP_NAME)
TOP_SRC_DIRS   = pkg
PACKAGES     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                  -exec dirname {} \\; | sort | uniq")
BINARY ?= go-hello-server

# This follows the output format for goreleaser
BINARY_LINUX_64 = ./dist/linux_amd64/go-hello-server

RELEASE_TAG ?= $(CIRCLE_TAG)
DOCKER_LATEST_TAG = darahayes/$(APP_NAME):latest
DOCKER_MASTER_TAG = darahayes/$(APP_NAME):master
DOCKER_RELEASE_TAG = darahayes/$(APP_NAME):$(RELEASE_TAG)

.PHONY: setup
setup:
	dep ensure

.PHONY: test
test:
	@echo Running tests:
	go test -v -race -cover

.PHONY: test-cover
test_cover:
	GOCACHE=off $(foreach pkg,$(PACKAGES),\
	go test -coverprofile=coverage.out -covermode=count $(addprefix $(PKG)/,$(pkg)) || exit 1;\
	tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: build
build: setup
	go build -o $(BINARY) ./cmd/hello-server/hello-server.go

.PHONY: build_linux
build_linux: setup
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX_64) ./cmd/hello-server/hello-server.go

.PHONY: docker_build
docker_build: build_linux
	docker build -t $(DOCKER_LATEST_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_build_release
docker_build_release:
	docker build -t $(DOCKER_LATEST_TAG) -t $(DOCKER_RELEASE_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_build_master
docker_build_master:
	docker build -t $(DOCKER_MASTER_TAG) --build-arg BINARY=$(BINARY_LINUX_64) .

.PHONY: docker_push_release
docker_push_release:
	@docker login --username $(DOCKERHUB_USERNAME) --password $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_LATEST_TAG)
	docker push $(DOCKER_RELEASE_TAG)
	
.PHONY: docker_push_master
docker_push_master:
	@docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_MASTER_TAG)