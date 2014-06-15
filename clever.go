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

// API supports basic auth with an API key or bearer auth with a token
type Auth struct {
	APIKey, Token string
}

type Clever struct {
	Auth
	Url string
}

// Using the oauth Transport pattern with API keys
type BasicAuthTransport struct {
	Username string
	Password string
}

func (bat BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
			bat.Username, bat.Password)))))
	return http.DefaultTransport.RoundTrip(req)
}

func (bat *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: bat}
}

// Creates a new clever object to make requests with. URL must be a valid base url, e.g. "https://api.clever.com"
func New(auth Auth, url string) *Clever {
	return &Clever{auth, url}
}

type CleverError struct {
	Code    string
	Message string `json:"error"`
}

func (err *CleverError) Error() string {
	if err.Code == "" {
		return err.Message
	}
	return fmt.Sprintf("%s (%s)", err.Error, err.Code)
}

type Paging struct {
	Count   int
	Current int
	Total   int
}

type Link struct {
	Rel string
	Uri string
}

type DistrictResp struct {
	District District `json:"data"`
	Links    []Link
	Uri      string
}

type District struct {
	Id   string
	Name string
}

type SchoolResp struct {
	Links  []Link
	School School `json:"data"`
	Uri    string
}

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

type TeacherResp struct {
	Links   []Link
	Teacher Teacher `json:"data"`
	Uri     string
}

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

type StudentResp struct {
	Links   []Link
	Student Student `json:"data"`
	Uri     string
}

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

type SectionResp struct {
	Links   []Link
	Section Section `json:"data"`
	Uri     string
}

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

type Location struct {
	Address string
	City    string
	State   string
	Zip     string
}

type Name struct {
	First  string
	Middle string
	Last   string
}

type Term struct {
	Name      string
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (clever *Clever) Query(path string, params url.Values, resp interface{}) error {
	uri := fmt.Sprintf("%s%s", clever.Url, path)
	if params != nil {
		uri = fmt.Sprintf("%s%s?%s", clever.Url, path, params.Encode())
	}

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
	if r.StatusCode != 200 {
		var error CleverError
		json.NewDecoder(r.Body).Decode(&error)
		return &error
	}
	err = json.NewDecoder(r.Body).Decode(resp)
	return err
}

type PagedResult struct {
	clever         Clever
	nextPagePath   string
	lastData       []map[string]interface{}
	lastDataCursor int
	lastError      error
}

func (clever *Clever) QueryAll(path string, params url.Values) PagedResult {
	paramString := ""
	if params != nil {
		paramString = "?" + params.Encode()
	}

	return PagedResult{clever: *clever, nextPagePath: path + paramString, lastDataCursor: -1}
}

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

func (r *PagedResult) Scan(result interface{}) error {
	data, err := json.Marshal(r.lastData[r.lastDataCursor]["data"])
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

func (r *PagedResult) Error() error {
	return r.lastError
}
