include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: test $(PKGS)
VERSION := $(shell cat VERSION)
SHELL := /bin/bash
PKGS := $(shell go list ./...)
$(eval $(call golang-version-check,1.8))

test: $(PKGS)
$(PKGS): golang-test-all-deps
	go get -t $@
	$(call golang-test-all,$@)

version.go:
	echo -e 'package clever\n\nconst Version = "$(VERSION)"' > version.go
