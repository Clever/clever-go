VERSION := $(shell cat VERSION)
SHELL := /bin/bash
PKG = github.com/Clever/clever-go
SUBPKGS =
PKGS = $(PKG) $(SUBPKGS)

.PHONY: test $(PKGS)

test: $(PKG)

version.go:
	echo -e 'package clever\n\nconst Version = "$(VERSION)"' > version.go

$(PKG):
ifeq ($(LINT),1)
	golint $(GOPATH)/src/$@*/**.go
endif
	go get -d -t $@
ifeq ($(COVERAGE),1)
	go test -cover -coverprofile=$(GOPATH)/src/$@/c.out $@ -test.v
	go tool cover -html=$(GOPATH)/src/$@/c.out
else
	go test $@ -test.v
endif
