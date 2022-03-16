########################################################################################################################
# Copyright (c) 2019 meshbox Foundation
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOPATH=$(shell go env GOPATH)

BUILD_TARGET_OS=$(shell go env GOOS)
BUILD_TARGET_ARCH=$(shell go env GOARCH)

# Pkgs

ROOT_PKG := "github.com/smartpom/pomgo"

default: clean build

build: pom alternate data simple struct

.PHONY: build

pom:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/pom -v ./cmd/pom

alternate:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/alternate -v ./cmd/alternate

data:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/data -v ./cmd/data

simple:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/simple -v ./cmd/simple

struct:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./bin/struct -v ./cmd/struct

.PHONY: clean
clean:
	@echo "Cleaning..."
	$(ECHO_V)rm -rf ./bin/*

