APP_NAME := user
APP_VERSION := $(shell git describe --tags --always)

.PHONY: dev
dev:
	@CONF_FILE_PATH=${PWD}/config/dev.json go run main.go

APP_PKG := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")
BUILD_TIME := $(shell date -u +"%FT%TZ")
GIT_REVISION := $(shell git rev-parse HEAD)

.PHONY: build
build:
	@mkdir -p ./bin
	GO111MODULE=on go build -ldflags "-X ${APP_PKG}/src/util.Name=${APP_NAME} \
	-X ${APP_PKG}/src/util.Version=${APP_VERSION} \
	-X ${APP_PKG}/src/util.BuildTime=${BUILD_TIME} \
	-X ${APP_PKG}/src/util.GitRevision=${GIT_REVISION}" \
	-o ./bin/user main.go

DOCKER_IMAGE_TAG := ${APP_NAME}:${APP_VERSION}
.PHONY: image
image:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build
	docker build --rm -t ${APP_NAME}:${APP_VERSION} .
