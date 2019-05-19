TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
PRACTICE_DIR = $(TOP_DIR)/practice
C_DIR = $(PRACTICE_DIR)/c

.PHONY: all
all: format build test

.PHONY: build
build: $(patsubst %.c, %.o, $(wildcard $(C_DIR)/*.c))

.PHONY: clean
clean:
	rm -f $(C_DIR)/*.o

.PHONY: format
format:
	@find -name '*.go' | xargs -n 1 go fmt

.PHONY: tags
tags:
	gotags -R main.go cc > tags

.PHONY: test
test: test-go
	@find $(PRACTICE_DIR) -name test.sh | sort | while read f; do \
	  echo $$f; \
	  $$f; \
	  echo; \
	done

.PHONY: test-go
test-go:
	@find ./practice -name cc | sort | xargs go test
