package test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Clever/clever-go/client"
	"github.com/Clever/clever-go/client/sections"
	"github.com/Clever/clever-go/client/users"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"

	httptransport "github.com/go-openapi/runtime/client"
)

type SyncTestSuite struct {
	suite.Suite
	mockClever   *httptest.Server
	cleverClient *client.Clever
	auth         runtime.ClientAuthInfoWriter
}

func (s *SyncTestSuite) SetupTest() {
	assert := s.Suite.Assert()

	// setup mock api server
	handler := func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/sections":
			sections, err := ioutil.ReadFile("mocks/sections.json")
			s.Suite.Assert().Nil(err)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Write(sections)
		case "/users":
			users, err := ioutil.ReadFile("mocks/users.json")
			s.Suite.Assert().Nil(err)
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.Write(users)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	s.mockClever = server

	// setup client
	u, err := url.Parse(s.mockClever.URL)
	assert.Nil(err)
	host := u.Host
	base := ""
	schemes := []string{"http"}

	cfg := client.DefaultTransportConfig().WithHost(host).WithBasePath(base).WithSchemes(schemes)
	c := client.NewHTTPClientWithConfig(strfmt.NewFormats(), cfg)
	assert.Nil(err)
	assert.IsType(c, &client.Clever{})
	s.cleverClient = c

	// setup auth
	s.auth = httptransport.BearerToken("TEST_TOKEN")
}

func (s *SyncTestSuite) TearDownSuite() {
	s.mockClever.Close()
}

// TestGetSections GET API /sections
func (s *SyncTestSuite) TestGetSections() {
	assert := s.Suite.Assert()

	ctx := context.Background()
	limit := int64(10000)
	params := &sections.GetSectionsParams{
		Limit:   &limit,
		Context: ctx,
	}

	ok, err := s.cleverClient.Sections.GetSections(params, s.auth)
	assert.Nil(err)
	assert.NotNil(ok)
	data := ok.GetPayload().Data
	assert.NotNil(data)
	assert.Equal("5a0c91b95c7cc000010d417c", data[0].Data.District)

}

// TestGetUsers GET API /users
func (s *SyncTestSuite) TestGetUsers() {
	assert := s.Suite.Assert()

	ctx := context.Background()
	limit := int64(10000)
	params := &users.GetUsersParams{
		Limit:   &limit,
		Context: ctx,
	}

	ok, err := s.cleverClient.Users.GetUsers(params, s.auth)
	assert.Nil(err)
	assert.NotNil(ok)
	data := ok.GetPayload().Data
	assert.NotNil(data)
	assert.Equal("5a0c91b95c7cc000010d417c", data[0].Data.District)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSyncTestSuite(t *testing.T) {
	suite.Run(t, new(SyncTestSuite))
}
