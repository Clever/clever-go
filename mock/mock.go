package clever

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"github.com/ant0ine/go-urlrouter"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
)

type PostHandler func(*http.Request, map[string]string) error

// Loads a directory with json files representing mock resources. See ./data for an example
func NewMock(postHandler PostHandler, dir string, lastRequestHeader ...*map[string][]string) (*http.Client, string) {
	router := urlrouter.Router{
		Routes: []urlrouter.Route{
			urlrouter.Route{
				PathExp: "/v1.1/districts",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/districts.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/districts/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/districts.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/schools",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/schools.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/schools/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/schools.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/teachers",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/teachers.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/teachers/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/teachers.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/students",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/students.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/students/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/students.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/sections",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/sections.json", dir), fmt.Sprintf("%s/sections2.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/sections/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/sections.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/events",
				Dest:    MockResource(postHandler, fmt.Sprintf("%s/events.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/events/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/events.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/mock/rate/limiter",
				Dest:    MockResourceRateLimit(),
			},
			urlrouter.Route{
				PathExp: "/mock/error",
				Dest:    MockError(),
			},
		},
	}

	if err := router.Start(); err != nil {
		panic(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("mock server:", r.URL)
		route, params := router.FindRouteFromURL(r.URL)
		handler := route.Dest.(func(http.ResponseWriter, *http.Request, map[string]string))
		if len(lastRequestHeader) > 0 {
			*(lastRequestHeader[0]) = r.Header
		}
		handler(w, r, params)
	}))

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: "doesntmatter"},
	}

	return t.Client(), ts.URL
}

func MockResource(postHandler PostHandler, filenames ...string) func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		switch req.Method {
		case "POST":
			err := postHandler(req, params)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("{}")) // Don't care what response is because we're not testing the API response
			return
		case "GET":
			page := 1
			u, _ := url.Parse(req.RequestURI)
			query, _ := url.ParseQuery(u.RawQuery)
			if _, ok := query["page"]; ok {
				page, _ = strconv.Atoi(query["page"][0])
			}
			file, err := os.Open(filenames[page-1])
			if err != nil {
				http.Error(w, fmt.Sprintf("couldn't read %s", filenames[page-1]), http.StatusInternalServerError)
				return
			}
			io.Copy(w, file)
		default:
			http.Error(w, fmt.Sprintf("method not supported"), http.StatusMethodNotAllowed)
		}

	}
}

func MockResourceId(filename string) func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		file, err := os.Open(filename)
		if err != nil {
			http.Error(w, fmt.Sprintf("couldn't read %s", filename), http.StatusInternalServerError)
			return
		}
		decoder := json.NewDecoder(file)
		var data map[string]interface{}
		if err = decoder.Decode(&data); err != nil {
			http.Error(w, fmt.Sprintf("couldn't decode %s", filename), http.StatusInternalServerError)
			return
		}
		for _, obj := range data["data"].([]interface{}) {
			o := obj.(map[string]interface{})["data"].(map[string]interface{})
			if o["id"] == params["id"] {
				jsonResponse := make(map[string]interface{})
				jsonResponse["data"] = o
				enc := json.NewEncoder(w)
				enc.Encode(jsonResponse)
				return
			}
		}
		http.Error(w, "", http.StatusNotFound)
	}
}

func MockResourceRateLimit() func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		const statusTooManyRequests = 429
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Bucket"), "testbucket_1")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Limit"), "200")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Reset"), "1394506274")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Remaining"), "0")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Bucket"), "testbucket_2")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Limit"), "1200")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Reset"), "never!!")
		w.Header().Add(http.CanonicalHeaderKey("X-Ratelimit-Remaining"), "0")
		http.Error(w, "", statusTooManyRequests)
	}
}

func MockError() func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		http.Error(w, `{"code":1337,"error":"there was an error"}`, 500)
	}
}
