TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
PRACTICE_DIR = $(TOP_DIR)/practice

.PHONY: all
all: test

fmt:
	@find -name '*.go' | xargs -n 1 go fmt

test:
	@find ./practice -name cc | sort | xargs go test
	@find $(PRACTICE_DIR) -name test.sh | sort | while read f; do \
	  echo $$f; \
	  $$f; \
	  echo; \
	done
