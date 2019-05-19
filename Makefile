TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build

.PHONY: all
all: tags format test

.PHONY: format
format:
	@go fmt *.go
	@find ./cc -name '*.go' | xargs go fmt

.PHONY: tags
tags:
	gotags -R main.go cc > tags

.PHONY: test
test:
	@go test ./cc


PRACTICE_DIR = $(TOP_DIR)/practice
include $(PRACTICE_DIR)/config.mk
