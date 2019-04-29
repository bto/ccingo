TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
PRACTICE_DIR = $(TOP_DIR)/practice

.PHONY: all
all: test

.PHONY: fmt
fmt:
	@find -name '*.go' | xargs -n 1 go fmt

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
