include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: test $(PKGS)
VERSION := $(shell cat VERSION)
SHELL := /bin/bash
PKGS := $(shell go list ./...)
$(eval $(call golang-version-check,1.10))

test: $(PKGS)
$(PKGS): golang-test-all-deps
	go get -t $@
	$(call golang-test-all,$@)

version.go:
	echo -e 'package clever\n\nconst Version = "$(VERSION)"' > version.go


install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)