# NOTICE
This library is no longer maintained by Clever, but if you would like to generate your own using the swagger file, it is available at https://github.com/Clever/swagger-api.

# clever-go

clever-go is a Go library for the [Clever API](https://clever.com/developers/docs).

## Documentation
[![Go Reference](https://pkg.go.dev/badge/github.com/Clever/clever-go.svg)](https://pkg.go.dev/github.com/Clever/clever-go)

## Usage

Client
```go
import (
    "github.com/Clever/clever-go/client"
	"github.com/Clever/go-clever/client/sections"
	"github.com/go-openapi/strfmt"

	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
    cfg := client.DefaultTransportConfig()

	c := client.NewHTTPClientWithConfig(strfmt.NewFormats(), cfg)

	sectionsParams := sections.NewGetSectionsParams()
	auth := httptransport.BearerToken("BEARER_TOKEN")

	sections, err := c.Sections.GetSections(sectionsParams, auth)
}
```

Auth
```go
import (
    "github.com/Clever/clever-go/auth"
)

func main() {
    ctx := context.Background()

	tokens, err := GetTokens(ctx, "CLIENT_ID", "CLIENT_SECRET")
}
```

## Releasing a new version

`clever-go` is versioned with `go mod`.
To release a new version of the library, you should increment the version in the `VERSION` file according to [the semver spec](http://semver.org/) and create a tag with the corresponding version.

You can use [gitsem](https://github.com/clever/gitsem) to accomplish this all with one command.
For example, to release a new minor version, you can just run `gitsem minor && git push && git push --tag` from the repository.

## Developing

clever-go is generated using go-swagger tools:
```
make generate
```
for help:
```
make help
```
