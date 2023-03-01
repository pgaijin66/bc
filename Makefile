CUR_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
PATH_TO_BIN := ${CUR_DIR}/src
LIB_PATH := /usr/local/bin
FILE_PERM := 755
APP_NAME := bc

UNAME := $(shell uname)
ifeq ($(UNAME), Linux)
	SHELL := /bin/bash
endif
ifeq ($(UNAME), Darwin)
	SHELL := /bin/zsh
endif

.DEFAULT_GOAL := help
help: ## List targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.SILENT: install
install: ## Installs binary to standard library PATH4
	@rm -rf "${LIB_PATH}/${APP_NAME}"
	@cp "${PATH_TO_BIN}/${APP_NAME}" "${LIB_PATH}"
	@chmod "${FILE_PERM}" "${LIB_PATH}/${APP_NAME}"
	@source ~/.zshrc

.PHONY: uninstall
uninstall: ## Uninstalls application
	@rm -rf "${LIB_PATH}/${APP_NAME}"
	@source ~/.zshrc


.PHONY: tag-n-release
tag-n-release: ## Tags
	@git tag -a v1.0.1
	@git push --tags

