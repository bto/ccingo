TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build
PRACTICE_DIR = $(TOP_DIR)/practice

.PHONY: all
all: test

test:
	@for f in $(PRACTICE_DIR)/*.sh; do \
	  echo $$(basename $$f); \
	  $$f; \
	  echo; \
	done
