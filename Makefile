.PHONY: test get
SHELL=bash

test: get
	go test -i
	go test

get:
	go get -d ./
