package clever

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const debug = false

// Auth holds credentials for access to the API: basic auth with an API key or bearer auth with a token
type Auth struct {
	APIKey, Token string
}

// Clever wraps the Clever API at the specified URL e.g. "https://api.clever.com"
type Clever struct {
	Auth
	Url string
}

// BasicAuthTransport contains a user's auth credentials. Clever uses the OAuth Transport pattern with API keys.
type BasicAuthTransport struct {
	Username string
	Password string
}

// RoundTrip makes a request and returns the response
func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
			bat.Username, bat.Password)))))
	return http.DefaultTransport.RoundTrip(req)
}

// Client returns a new Client object for the specified BasicAuthTransport
func (bat *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: bat}
}

// New returns a new Clever object to make requests with. URL must be a valid base url, e.g. "https://api.clever.com"
func New(auth Auth, url string) *Clever {
	return &Clever{auth, url}
}

// CleverError contains an error that occurred within the Clever API
type CleverError struct {
	Code    string
	Message string `json:"error"`
}

// Error returns a string representation of a CleverError
func (err *CleverError) Error() string {
	if err.Code == "" {
		return err.Message
	}
	return fmt.Sprintf("%s (%s)", err.Error, err.Code)
}

// TooManyRequestsError indicates the number of requests has exceeded the rate limit
type TooManyRequestsError struct {
	Header http.Header
}

// TooManyRequestsError creates a TooManyRequestsError
func (err *TooManyRequestsError) Error() string {
	err_string := "Too Many Requests"
	for bucket_index, bucket_name := range err.Header[http.CanonicalHeaderKey("X-Ratelimit-Bucket")] {
		err_string += fmt.Sprintf("\nBucket: %s", bucket_name)
		for _, prop := range []string{"Remaining", "Limit", "Reset"} {
			headers_for_prop := err.Header[http.CanonicalHeaderKey("X-Ratelimit-"+prop)]
			if bucket_index < len(headers_for_prop) {
				err_string += fmt.Sprintf(", %s: %s", prop, headers_for_prop[bucket_index])
			}
		}
	}
	return err_string
}

// Paging contains information for paging a response
type Paging struct {
	Count   int
	Current int
	Total   int
}

// Link represents a stable link for querying the API
type Link struct {
	Rel string
	Uri string
}

// DistrictResp wraps the response given when the user queries for a District
type DistrictResp struct {
	District District `json:"data"`
	Links    []Link
	Uri      string
}

// District corresponds to the District resource in the Clever data schema: clever.com/schema
type District struct {
	Id   string
	Name string
}

// SchoolResp wraps the response given when the user queries for a School
type SchoolResp struct {
	Links  []Link
	School School `json:"data"`
	Uri    string
}

// School corresponds to the School resource in the Clever data schema: clever.com/schema
type School struct {
	Created      string
	District     string
	HighGrade    string `json:"high_grade"`
	Id           string
	LastModified string `json:"last_modified"`
	Location     Location
	LowGrade     string `json:"low_grade"`
	Name         string
	NcesId       string `json:"nces_id"`
	Phone        string
	SchoolNumber string `json:"school_number"`
	SisId        string `json:"sis_id"`
	StateId      string `json:"state_id"`
}

// TeacherResp wraps the response given when the user queries for a Teacher
type TeacherResp struct {
	Links   []Link
	Teacher Teacher `json:"data"`
	Uri     string
}

// Teacher corresponds to the Teacher resource in the Clever data schema: clever.com/schema
type Teacher struct {
	Created       string
	District      string
	Email         string
	Id            string
	LastModified  string `json:"last_modified"`
	Name          Name
	School        string
	SisId         string `json:"sis_id"`
	TeacherNumber string `json:"teacher_number"`
	Title         string
}

// StudentResp wraps the response given when the user queries for a Student
type StudentResp struct {
	Links   []Link
	Student Student `json:"data"`
	Uri     string
}

// Student corresponds to the Student resource in the Clever data schema: clever.com/schema
type Student struct {
	Created           string
	District          string
	Dob               string
	Email             string
	FrlStatus         string `json:"frl_status"`
	Gender            string
	Grade             string
	HispanicEthnicity string `json:"hispanic_ethnicity"`
	Id                string
	LastModified      string `json:"last_modified"`
	Location          Location
	Name              Name
	Race              string
	School            string
	SisId             string `json:"sis_id"`
	StateId           string `json:"state_id"`
	StudentNumber     string `json:"student_number"`
}

// SectionResp wraps the response given when the user queries for a Section
type SectionResp struct {
	Links   []Link
	Section Section `json:"data"`
	Uri     string
}

// Section corresponds to the Section resource in the Clever data schema: clever.com/schema
type Section struct {
	CourseName   string `json:"course_name"`
	CourseNumber string `json:"course_number"`
	Created      string
	District     string
	Grade        string
	Id           string `json:"id"`
	LastModified string `json:"last_modified"`
	Name         string
	School       string
	SisId        string `json:"sis_id"`
	Students     []string
	Subject      string
	Teacher      string
	Term
}

// Location represents a complete address for use with the Student and School resources
type Location struct {
	Address string
	City    string
	State   string
	Zip     string
}

// Name represents the full name of a Student or Teacher resource
type Name struct {
	First  string
	Middle string
	Last   string
}

// Term holds information about the duration of a school term
type Term struct {
	Name      string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

/*
Query makes a request to Clever given a Clever object, endpoint path, and parameters (pass in nil for no parameters).
*/
func (clever *Clever) Query(path string, params url.Values, resp interface{}) error {
	// Create request URI from Clever base, path, params
	uri := fmt.Sprintf("%s%s", clever.Url, path)
	if params != nil {
		uri = fmt.Sprintf("%s%s?%s", clever.Url, path, params.Encode())
	}

	// Ensure authentication is provided
	var client *http.Client
	req, _ := http.NewRequest("GET", uri, nil)
	if clever.Auth.Token != "" {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: clever.Auth.Token},
		}
		client = t.Client()
	} else if clever.Auth.APIKey != "" {
		t := &BasicAuthTransport{
			Username: clever.Auth.APIKey,
		}
		client = t.Client()
	} else {
		return fmt.Errorf("Must provide either API key or bearer token")
	}
	if debug {
		log.Printf("get { %v } -> {\n", uri)
	}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}
	if r.StatusCode == 429 {
		return &TooManyRequestsError{r.Header}
	} else if r.StatusCode != 200 {
		var error CleverError
		json.NewDecoder(r.Body).Decode(&error)
		return &error
	}
	err = json.NewDecoder(r.Body).Decode(resp)
	return err
}

// PagedResult wraps a response. It allows for paged reading of a response in conjunction with QueryAll() and Next()
type PagedResult struct {
	clever         Clever
	nextPagePath   string
	lastData       []map[string]interface{}
	lastDataCursor int
	lastError      error
}

// QueryAll returns a PagedResult to allow for paged reading of large responses from the Clever API
func (clever *Clever) QueryAll(path string, params url.Values) PagedResult {
	paramString := ""
	if params != nil {
		paramString = "?" + params.Encode()
	}

	return PagedResult{clever: *clever, nextPagePath: path + paramString, lastDataCursor: -1}
}

/*
Next returns true if a PagedResult contains additional data and false if the cursor has reached
the end of the available data for this response.
*/
func (r *PagedResult) Next() bool {
	if r.lastDataCursor != -1 && r.lastDataCursor < len(r.lastData)-1 {
		r.lastDataCursor++
		return true
	}

	if r.nextPagePath == "" {
		return false
	}

	resp := &struct {
		Data   []map[string]interface{}
		Links  []Link
		Paging Paging
	}{}
	r.lastError = r.clever.Query(r.nextPagePath, nil, resp)
	if r.lastError != nil {
		return false
	}
	r.lastData = resp.Data
	r.lastDataCursor = 0
	r.nextPagePath = ""
	for _, link := range resp.Links {
		if link.Rel == "next" {
			r.nextPagePath = link.Uri
			break
		}
	}
	return len(r.lastData) > 0
}

// Scan parses the next page of results in a PagedResult r into result. Scan throws an error if r is invalid JSON.
func (r *PagedResult) Scan(result interface{}) error {
	data, err := json.Marshal(r.lastData[r.lastDataCursor]["data"])
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Error returns the error in a response, if there is one
func (r *PagedResult) Error() error {
	return r.lastError
}
