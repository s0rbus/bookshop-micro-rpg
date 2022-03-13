
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
ARCH=amd64
LINUX=$(BINARY)-linux-$(ARCH)
WINDOWS=$(BINARY)-windows-$(ARCH).exe
DARWIN=$(BINARY)-darwin-$(ARCH)

#Note, when not using main package for variable setting using LDFLAGS, have to provide full package path, not just package name
LDFLAGS=-ldflags="-X main.version=$(VERSION) -X main.buildstamp=$(BUILDTS) -X main.githash=$(GITSTATUS)"

all: windows linux darwin
	$(shell mkdir -p expansions && find expansion-src -name '*.so' -exec cp {} expansions/ \;)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(LINUX): $(SRC) ${ROOT}/CHANGELOG.md $(ESO)
	# @go build -trimpath -o $(BINARY) $(LDFLAGS) $(SRC)
	env GOOS=linux GOARCH=$(ARCH) go build -trimpath -o $(LINUX) $(LDFLAGS) $(SRC)

$(WINDOWS): $(SRC) ${ROOT}/CHANGELOG.md
	env GOOS=windows GOARCH=$(ARCH) go build -trimpath -o $(WINDOWS) $(LDFLAGS) $(SRC)
	# env GOOS=windows GOARCH=386 go build -trimpath -o $(WINDOWS) $(LDFLAGS) $(SRC)

$(DARWIN): $(SRC) ${ROOT}/CHANGELOG.md
	env GOOS=darwin GOARCH=$(ARCH) go build -trimpath -o $(DARWIN) $(LDFLAGS) $(SRC)

$(ESO): $(ESRC)
	@./build-expansions.sh

expansions: $(ESO)

.PHONY: clean

clean: 
	$(shell rm -f $(BINARY)-*)
	$(shell find expansion-src -name '*.so' -exec rm {} \;)
	$(shell find expansions -name '*.so' -exec rm {} \;)



