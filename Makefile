# Project output directory.
OUTPUT_DIR := ./bin

# Module name.
NAME := vkectl

# This repo's root import path (under GOPATH).
ROOT := github.com/volcengine/vkectl

# Current version of the project.
VERSION      ?= $(shell git describe --tags --always --dirty)
BRANCH       ?= $(shell git branch | grep \* | cut -d ' ' -f2)
GITCOMMIT    ?= $(shell git rev-parse HEAD)
GITTREESTATE ?= $(if $(shell git status --porcelain),dirty,clean)
BUILDDATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Project main package location.
CMD_DIR := ./main

all: build

build:
	@go build -v -o $(OUTPUT_DIR)/$(NAME)                                  \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);
