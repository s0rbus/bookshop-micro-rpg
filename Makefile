NAME=$(shell basename $(CURDIR))

SHELL=/bin/bash

# The name of the executable (default is current directory name)
BINARY := $(NAME)
ROOT := $(shell git rev-parse --show-toplevel)
#Extract version number from CHANGELOG file. have to escape '#' otherwise make sees it as comment
VERSION := $(shell sed -n 's/^\#\# \[\([0-9]\.[0-9]\.[0-9].*\)\].*/\1/p' ${ROOT}/CHANGELOG.md | head -1)
BUILDTS := $(shell date '+%Y%m%d_%H:%M:%S')
GITHASH := $(shell git rev-parse --short HEAD)
GITSTATUS := "$(GITHASH)$(shell git diff --quiet || echo '-dirty')"
SRC := $(wildcard *.go)
subdirs := $(wildcard expansions/*/)
ESRC := $(wildcard $(addsuffix *.go,$(subdirs)))
ESO := $(patsubst %.go,%.so,$(ESRC))

#Note, when not using main package for variable setting using LDFLAGS, have to provide full package path, not just package name
LDFLAGS=-ldflags="-X main.version=$(VERSION) -X main.buildstamp=$(BUILDTS) -X main.githash=$(GITSTATUS)"

bin/$(BINARY): $(SRC) ${ROOT}/CHANGELOG.md
	@go build -trimpath -o bin/$(BINARY) $(LDFLAGS) $(SRC)

$(ESO): $(ESRC)
	@./build-expansions.sh

all: bin/$(BINARY) $(ESO)
	@./make-dist.sh

expansions: $(ESO)

#dist: $(ESO)
#	@./make-dist.sh

.PHONY: clean

clean: 
	@rm -rf bin
	@$(shell find expansions -name '*.so' -exec rm {} \;)


