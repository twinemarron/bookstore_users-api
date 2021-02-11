# Go パラメータ
GOCMD=go
GOBUILD=$(GOCMD) build
# GOCLEAN=$(GOCMD) clean
# GOTEST=$(GOCMD) test
# GOGET=$(GOCMD) get
BINARY_DIR=./bin/
BINARY_NAME=mybinary
# BINARY_UNIX=$(BINARY_NAME)_unix
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")

# all: test build
build:
	$(GOBUILD) -o $(BINARY_DIR)$(BINARY_NAME) -v

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: imports
imports:
	goimports -w $(GOFILES)
