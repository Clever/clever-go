package test

import (
	"os"
	"testing"

	"github.com/Clever/clever-go/client"
	"github.com/Clever/clever-go/client/sections"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"

	httptransport "github.com/go-openapi/runtime/client"
)

// TestGetSections API Client
// - set the CLEVER_API_TOKEN env var
func TestGetSections(t *testing.T) {
	assert := assert.New(t)

	cfg := client.DefaultTransportConfig()

	c := client.NewHTTPClientWithConfig(strfmt.NewFormats(), cfg)

	assert.IsType(c, &client.Clever{})

	sectionsParams := sections.NewGetSectionsParams()
	auth := httptransport.BearerToken(getTokenFromEnv(assert))

	sections, err := c.Sections.GetSections(sectionsParams, auth)

	assert.Nil(err)
	assert.NotNil(sections)
}

func getTokenFromEnv(assert *assert.Assertions) string {
	token := os.Getenv("CLEVER_API_TOKEN")
	if token == "" {
		assert.FailNow("Please set the CLEVER_API_TOKEN env var.")
	}
	return token
}
