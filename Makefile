.PHONY: dev
dev:
	@CONF_FILE_PATH=${PWD}/config/dev.json go run main.go

APP_NAME := user
APP_VERSION := $(shell git describe --tags --always)
APP_PKG := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

.PHONY: image
image:
	docker build \
	--build-arg APP_NAME=${APP_NAME} \
	--build-arg APP_VERSION=${APP_VERSION} \
	--build-arg APP_PKG=${APP_PKG} \
	-t ${APP_NAME}:${APP_VERSION} .
