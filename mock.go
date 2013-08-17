package clever

import (
	"encoding/json"
	"fmt"
	"github.com/ant0ine/go-urlrouter"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

// Loads a directory with json files representing mock resources. See ./data for an example
func NewMock(dir string) *Clever {
	router := urlrouter.Router{
		Routes: []urlrouter.Route{
			urlrouter.Route{
				PathExp: "/v1.1/districts",
				Dest:    MockResource(fmt.Sprintf("%s/districts.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/districts/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/districts.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/schools",
				Dest:    MockResource(fmt.Sprintf("%s/schools.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/schools/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/schools.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/teachers",
				Dest:    MockResource(fmt.Sprintf("%s/teachers.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/teachers/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/teachers.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/students",
				Dest:    MockResource(fmt.Sprintf("%s/students.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/students/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/students.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/sections",
				Dest:    MockResource(fmt.Sprintf("%s/sections.json", dir)),
			},
			urlrouter.Route{
				PathExp: "/v1.1/sections/:id",
				Dest:    MockResourceId(fmt.Sprintf("%s/sections.json", dir)),
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
		handler(w, r, params)
	}))

	return &Clever{Auth{"doesntmatter", ""}, ts.URL}
}

func MockResource(filename string) func(http.ResponseWriter, *http.Request, map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		file, err := os.Open(filename)
		if err != nil {
			http.Error(w, fmt.Sprintf("couldn't read %s", filename), http.StatusInternalServerError)
			return
		}
		io.Copy(w, file)
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
