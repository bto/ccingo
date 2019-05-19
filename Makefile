TOP_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
BUILD_DIR = $(TOP_DIR)/build

.PHONY: all
all: tags

.PHONY: tags
tags:
	gotags -R main.go cc > tags

PRACTICE_DIR = $(TOP_DIR)/practice
include $(PRACTICE_DIR)/config.mk
