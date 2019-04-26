TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
PRACTICE_DIR = $(TOP_DIR)/practice

.PHONY: all
all: test

fmt:
	@find -name '*.go' | xargs -n 1 go fmt

test:
	@find $(PRACTICE_DIR) -name test.sh | while read f; do \
	  [ ! -x $$f ] && continue; \
	  echo $$(basename $$f); \
	  $$f; \
	  echo; \
	done
