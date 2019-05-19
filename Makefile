TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build

.PHONY: all
all: tags format test

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/*

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
	$(TOP_DIR)/test.sh


PRACTICE_DIR = $(TOP_DIR)/practice
include $(PRACTICE_DIR)/config.mk
