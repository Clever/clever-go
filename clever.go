package clever

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"runtime"
)

const debug = false

var unameCache []byte

// Clever wraps the Clever API at the specified URL e.g. "https://api.clever.com"
type Clever struct {
	client *http.Client
	url    string
}

// BasicAuthTransport is an http.Transport that performs HTTP Basic Auth.
type BasicAuthTransport struct {
	Username string
	Password string
}

func setTrackingHeaders(req *http.Request) {
	// Want to swallow errors here because desired behavior is nil --> blank string if value doesn't exist
	if unameCache == nil {
		unameCache, _ = exec.Command("uname", "-a").Output()
	}
	customUA, _ := json.Marshal(map[string]string{
		"lang":             "go",
		"lang_version":     runtime.Version(),
		"platform":         runtime.GOOS,
		"uname":            fmt.Sprintf("%s", unameCache),
		"publisher":        "clever",
		"bindings_version": Version,
	})
	req.Header.Add("User-Agent", "Clever/GoBindings/"+Version)
	req.Header.Add("X-Clever-Client-User-Agent", string(customUA))
}

// RoundTrip makes a request and returns the response
func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
			bat.Username, bat.Password)))))
	setTrackingHeaders(req)
	return http.DefaultTransport.RoundTrip(req)
}

// Client returns a new Client object for the specified BasicAuthTransport
func (bat *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: bat}
}

// New returns a new Clever object to make requests with. URL must be a valid base url, e.g. "https://api.clever.com"
func New(client *http.Client, url string) *Clever {
	return &Clever{client, url}
}

// CleverError contains an error that occurred within the Clever API
type CleverError struct {
	Code    int
	Message string `json:"error"`
}

// Error returns a string representation of a CleverError
func (err *CleverError) Error() string {
	if err.Code == 0 {
		return err.Message
	}
	return fmt.Sprintf("%s (%d)", err.Message, err.Code)
}

// TooManyRequestsError indicates the number of requests has exceeded the rate limit
type TooManyRequestsError struct {
	Header http.Header
}

// TooManyRequestsError creates a TooManyRequestsError
func (err *TooManyRequestsError) Error() string {
	errString := "Too Many Requests"
	for bucketIndex, bucketName := range err.Header[http.CanonicalHeaderKey("X-Ratelimit-Bucket")] {
		errString += fmt.Sprintf("\nBucket: %s", bucketName)
		for _, prop := range []string{"Remaining", "Limit", "Reset"} {
			headersForProp := err.Header[http.CanonicalHeaderKey("X-Ratelimit-"+prop)]
			if bucketIndex < len(headersForProp) {
				errString += fmt.Sprintf(", %s: %s", prop, headersForProp[bucketIndex])
			}
		}
	}
	return errString
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
	Id        string
	Name      string
	MdrNumber string `json:"mdr_number"`
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
	Credentials   Credentials
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
	Credentials	  Credentials
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
	Id           string
	LastModified string `json:"last_modified"`
	Name         string
	School       string
	SisId        string `json:"sis_id"`
	Students     []string
	Subject      string
	Teacher      string
	Teachers     []string
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

type Credentials struct {
	DistrictUsername string `json:"district_username"`
	DistrictPassword string `json:"district_password"`
}

// Query makes a GET request to Clever using the Request function (see below)
func (clever *Clever) Query(path string, params url.Values, resp interface{}) error {
	return clever.Request("GET", path, params, nil, resp)
}

// Request makes a request to Clever given a Clever object, http method,
// endpoint path, parameters (pass in nil for no parameters), and a body .
func (clever *Clever) Request(method string, path string, params url.Values, body interface{}, resp interface{}) error {

	// Create request URI from Clever base, path, params
	uri := fmt.Sprintf("%s%s", clever.url, path)
	if params != nil {
		uri = fmt.Sprintf("%s%s?%s", clever.url, path, params.Encode())
	}

	var bodyReader io.Reader
	if body != nil {
		rawBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(rawBody)
	}

	// Ensure authentication is provided
	req, _ := http.NewRequest(method, uri, bodyReader)
	req.Header.Add("Content-Type", "application/json")
	setTrackingHeaders(req)
	if debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Printf("request:\n")
		log.Printf("%v\n", string(dump))
	}
	r, err := clever.client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n", string(dump))
	}
	if r.StatusCode == 429 {
		return &TooManyRequestsError{r.Header}
	} else if r.StatusCode != 200 {
		var error CleverError
		if err := json.NewDecoder(r.Body).Decode(&error); err != nil {
			return err
		}
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

// Next returns true if a PagedResult contains additional data and false if the
// cursor has reached the end of the available data for this response.
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
