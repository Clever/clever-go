# clever-go

clever-go is a Go wrapper around the [Clever API](https://clever.com/developers/docs).

## Documentation

[![GoDoc](https://godoc.org/github.com/Clever/clever-go?status.png)](https://godoc.org/github.com/Clever/clever-go).

## Developing

clever-go is built and tested against Go 1.2.
Ensure this is the version of Go you're running with `go version`.
Make sure your GOPATH is set, e.g. `export GOPATH=~/go`.
Clone the repository to a location outside your GOPATH, and symlink it to `$GOPATH/src/github.com/Clever/clever-go`.
The makefile may require [Mercurial](http://mercurial.selenic.com/downloads) to run the first time.
If you have done all of the above, then you should be able to run

```
make
```
## Troubleshooting
Generating godocs requires the godoc command line tool, which is not included in the Go installation. Install godoc with

```
go install code.google.com/p/go.tools/cmd/godoc
```

You may need to add $GOROOT export to .bashrc and $PATH as per this [issue](https://github.com/Homebrew/homebrew/issues/23281#issuecomment-29672689):

    export GOROOT=`go env GOROOT`
    export PATH=$PATH:$GOROOT/bin
