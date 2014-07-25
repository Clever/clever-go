# clever-go

clever-go is a Go library for the [Clever API](https://clever.com/developers/docs).

## Documentation

[![GoDoc](https://godoc.org/github.com/Clever/clever-go?status.png)](https://godoc.org/github.com/Clever/clever-go).

## Developing

clever-go is built and tested against Go 1.3.
Ensure this is the version of Go you're running with `go version`.
Make sure your GOPATH is set, e.g. `export GOPATH=~/go`.
Clone the repository to a location outside your GOPATH, and symlink it to `$GOPATH/src/github.com/Clever/clever-go`.
Since some dependencies are Mercurial repositories, ensure that you have [Mercurial](http://mercurial.selenic.com/downloads)  installed.
The `godoc` command line tool can be used to generate documentation from comments in source. This is useful for enforcing good commenting practices and ensures a standard, readable format for the resulting documentation. If `godoc` is not already installed, instructions to install can be found [here](http://golang.org/doc/go1.2#go_tools_godoc).
If you have done all of the above, then you should be able to run

```
make
```
