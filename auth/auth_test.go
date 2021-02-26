package auth

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAuthTokens
// - set the CLEVER_API_CLIENT_ID and CLEVER_API_CLIENT_SECRET env var
func TestGetAuthTokens(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()

	tokens, err := GetTokens(ctx, getClientIDFromEnv(assert), getClientSecretFromEnv(assert))

	assert.Nil(err)
	assert.GreaterOrEqual(len(tokens.Data), 1)
}

func getClientIDFromEnv(assert *assert.Assertions) string {
	id := os.Getenv("CLEVER_API_CLIENT_ID")
	if id == "" {
		assert.FailNow("Please set the CLEVER_API_CLIENT_ID env var.")
	}
	return id
}

func getClientSecretFromEnv(assert *assert.Assertions) string {
	secret := os.Getenv("CLEVER_API_CLIENT_SECRET")
	if secret == "" {
		assert.FailNow("Please set the CLEVER_API_CLIENT_SECRET env var.")
	}
	return secret
}
