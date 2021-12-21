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

publish: clean build_mac build_linux_amd64 build_windows build_linux_arm64

build:
	@go build -v -o $(OUTPUT_DIR)/$(NAME)                                  \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);

generate:
	hack/gen_client/generate.sh
	go run hack/gen_client/main.go

clean:
	rm -f $(OUTPUT_DIR)/$(NAME)*

install: build
	cp $(OUTPUT_DIR)/$(NAME) /usr/local/bin

build_mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64                                \
	  go build -v -o $(OUTPUT_DIR)/$(NAME)                                 \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);
	tar zcvf $(OUTPUT_DIR)/$(NAME)-macosx-${VERSION}-amd64.tgz -C $(OUTPUT_DIR) $(NAME)

build_linux_amd64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64                                 \
	  go build -v -o $(OUTPUT_DIR)/$(NAME)                                 \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);
	tar zcvf $(OUTPUT_DIR)/$(NAME)-linux-${VERSION}-amd64.tgz -C $(OUTPUT_DIR) $(NAME)

build_windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64                               \
	  go build -v -o $(OUTPUT_DIR)/$(NAME).exe                             \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);
	cd $(OUTPUT_DIR); zip $(NAME)-windows-${VERSION}-amd64.zip $(NAME).exe; cd -

build_linux_arm64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64                                 \
	  go build -v -o $(OUTPUT_DIR)/$(NAME)                                 \
	  -ldflags "-s -w -X $(ROOT)/pkg/version.module=$(NAME)                \
	    -X $(ROOT)/pkg/version.version=$(VERSION)                          \
	    -X $(ROOT)/pkg/version.branch=$(BRANCH)                            \
	    -X $(ROOT)/pkg/version.gitCommit=$(GITCOMMIT)                      \
	    -X $(ROOT)/pkg/version.gitTreeState=$(GITTREESTATE)                \
	    -X $(ROOT)/pkg/version.buildDate=$(BUILDDATE)"                     \
	  $(CMD_DIR);
	tar zcvf $(OUTPUT_DIR)/$(NAME)-linux-${VERSION}-arm64.tgz -C $(OUTPUT_DIR) $(NAME)