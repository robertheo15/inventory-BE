SHELL := /bin/bash

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

BASE := $(shell pwd)
PKGS := $(shell go list ./... | grep -v /vendor/ | grep -v /internal/mock)
ALL_PACKAGES := $(shell go list ./... | grep -v /vendor/ | grep -v /internal/mock)
GOLANGCI_CMD := $(shell command -v golangci-lint 2> /dev/null)

.Phony: lint fmt check-golangci

fmt:
	@echo -e "$(OK_COLOR)==> formatting code$(NO_COLOR)..."
	@go fmt $(ALL_PACKAGES)

check-golangci:
ifndef GOLANGCI_CMD
    $(error "Please install golangci linters from https://golangci-lint.run/usage/install/")
endif

lint: fmt check-golangci
	@echo -e "$(OK_COLOR)==> linting projects$(NO_COLOR)..."
	@env golangci-lint run --fix
	@echo -e "$(OK_COLOR)==> all is good."