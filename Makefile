
SHELL=/bin/bash

# The name of the executable
BINARY := bookshop-micro-rpg
ROOT := $(shell git rev-parse --show-toplevel)
#Extract version number from CHANGELOG file. have to escape '#' otherwise make sees it as comment
VERSION := $(shell sed -n 's/^\#\# \[\([0-9]\.[0-9]\.[0-9].*\)\].*/\1/p' ${ROOT}/CHANGELOG.md | head -1)
BUILDTS := $(shell date '+%Y%m%d_%H:%M:%S')
GITHASH := $(shell git rev-parse --short HEAD)
GITSTATUS := "$(GITHASH)$(shell git diff --quiet || echo '-dirty')"
SRC := $(wildcard *.go)
subdirs := $(wildcard expansion-src/*/)
ESRC := $(wildcard $(addsuffix *.go,$(subdirs)))
ESO := $(patsubst %.go,%.so,$(ESRC))

#Note, when not using main package for variable setting using LDFLAGS, have to provide full package path, not just package name
LDFLAGS=-ldflags="-X main.version=$(VERSION) -X main.buildstamp=$(BUILDTS) -X main.githash=$(GITSTATUS)"

$(BINARY): $(SRC) ${ROOT}/CHANGELOG.md
	@go build -trimpath -o $(BINARY) $(LDFLAGS) $(SRC)

$(ESO): $(ESRC)
	@./build-expansions.sh

all: $(BINARY) $(ESO)
	@./make-dist.sh

expansions: $(ESO)

#dist: $(ESO)
#	@./make-dist.sh

.PHONY: clean

clean: 
	@rm -rf dist
	@rm -f $(BINARY)
	@$(shell find expansion-src -name '*.so' -exec rm {} \;)


