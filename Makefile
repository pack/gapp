.PHONY: all build install test

CURRENT_DIR = $(shell  cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

VERSION = $(shell cat $(CURRENT_DIR)/VERSION)
ifeq ($(strip $(shell git status --porcelain)),)
	GITCOMMIT = $(shell git rev-parse --short HEAD)
else
	GITCOMMIT = $(shell git rev-parse --short HEAD)-dirty
endif

LDFLAGS="-X main.GITCOMMIT '$(GITCOMMIT)' -X main.VERSION '$(VERSION)' -w"

default: all

all: build

build:
	@cd $(CURRENT_DIR)
	go build -ldflags $(LDFLAGS)

install:
	@cd $(CURRENT_DIR)
	go install -ldflags $(LDFLAGS)

test:
	go test
