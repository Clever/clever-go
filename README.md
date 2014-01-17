# clever-go

[![Build Status](https://secure.travis-ci.org/Clever/clever-go.png)](http://travis-ci.org/Clever/clever-go)

# Usage

```go
package main

import (
	"fmt"
	clevergo "github.com/Clever/clever-go"
	"os"
)

func main() {
	clever := clevergo.New(clevergo.Auth{"DEMO_KEY", ""}, "https://api.clever.com")
	resp := &clevergo.DistrictsResp{}
	if err := clever.Query("/v1.1/districts", map[string]string{}, resp); err != nil {
		fmt.Fprintf(os.Stderr, "Error getting districts: %s \n", err)
	}

	for _, dresp := range resp.Districts {
		fmt.Println(dresp.District)
	}
}
```
