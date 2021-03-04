package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	// DefaultHost is the default Host
	DefaultHost string = "clever.com"
	// DefaultBasePath is the default BasePath
	DefaultBasePath string = "/oauth/"
	// DefaultOwnerType is the default owner type
	// - district is the only supported value
	DefaultOwnerType string = "district"
)

// AuthenticationError represents an auth api call error
type AuthenticationError struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

// TokenResponse reprents the token response form an auth api call
type TokenResponse struct {
	Data []struct {
		ID      string    `json:"id"`
		Created time.Time `json:"created"`
		Owner   struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"owner"`
		AccessToken string   `json:"access_token"`
		Scopes      []string `json:"scopes"`
	} `json:"data"`
	Links []struct {
		Rel string `json:"rel"`
		URI string `json:"uri"`
	} `json:"links"`
}

// GetTokens - takes a client id and a client secret and returns bearer tokens
func GetTokens(ctx context.Context, clientID, clientSecret string) (*TokenResponse, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("error client id and client secret are required")
	}

	c := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getAuthURL(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", getAuthEncoding(clientID, clientSecret))

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making auth call: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, getErrorFromBody(resp.Body)
	}

	tr, err := getTokenResponseFromBody(resp.Body)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// GetDistrictTokens - takes a client id and a client secret and returns bearer tokens
func GetDistrictTokens(ctx context.Context, clientID, clientSecret, district string) (*TokenResponse, error) {
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("error client id and client secret are required")
	}

	c := &http.Client{}

	url := getAuthURL() + "&district=" + district
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", getAuthEncoding(clientID, clientSecret))

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making auth call: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, getErrorFromBody(resp.Body)
	}

	tr, err := getTokenResponseFromBody(resp.Body)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// getAuthURL - builds the auth url from the libraries constants
func getAuthURL() string {
	return "https://" + DefaultHost + DefaultBasePath + "tokens?owner_type=" + DefaultOwnerType
}

// getAuthEncoding - base64 encodes the client id and client secret as expected by the oauth api
func getAuthEncoding(clientID, clientSecret string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	return "Basic " + b64
}

// getErrorFromBody - decodes an authentication error from the response body
func getErrorFromBody(body io.ReadCloser) error {
	defer body.Close()

	authErr := &AuthenticationError{}
	err := json.NewDecoder(body).Decode(authErr)
	if err != nil {
		return fmt.Errorf("error decoding response body from auth api: %s", err)
	}

	return fmt.Errorf("authentication error: %s, description: %s", authErr.Error, authErr.Description)
}

// getTokenResponseFromBody - decodes an tokens response from the response body
func getTokenResponseFromBody(body io.ReadCloser) (*TokenResponse, error) {
	defer body.Close()

	tr := &TokenResponse{}
	err := json.NewDecoder(body).Decode(tr)
	if err != nil {
		return nil, fmt.Errorf("error decoding token response body from auth api: %s", err)
	}

	return tr, nil
}
