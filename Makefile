TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
C_DIR = $(TOP_DIR)/c

.PHONY: all
all: build tags format test

.PHONY: build
build: $(patsubst %.c, %.o, $(wildcard $(C_DIR)/*.c))

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
	@echo
	@$(TOP_DIR)/test.sh


PRACTICE_DIR = $(TOP_DIR)/practice
include $(PRACTICE_DIR)/config.mk
