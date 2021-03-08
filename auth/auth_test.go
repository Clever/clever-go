package auth

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	mockClever *httptest.Server
}

func (s *AuthTestSuite) SetupTest() {
	assert := s.Suite.Assert()

	// setup mock api server
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth/tokens":
			sections, err := ioutil.ReadFile("mocks/tokens.json")
			s.Suite.Assert().Nil(err)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Write(sections)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	s.mockClever = server

	u, err := url.Parse(server.URL)
	assert.Nil(err)
	DefaultHost = u.Host
	defaultProto = "http"
}

func (s *AuthTestSuite) TearDownSuite() {
	s.mockClever.Close()
}

// TestGetTokens GET API /tokens
func (s *AuthTestSuite) TestGetTokens() {
	assert := s.Suite.Assert()

	ctx := context.Background()
	tokens, err := GetTokens(ctx, "CLIENT_ID", "CLIENT_SECRET")

	assert.Nil(err)
	assert.NotNil(tokens)
	assert.Equal("12345fcabb18e535b594ed2096c5c8e242362be5", tokens.Data[0].AccessToken)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
