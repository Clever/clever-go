package clever

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	mock "github.com/Clever/clever-go/mock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var dummytransport = &oauth.Transport{
	Token: &oauth.Token{AccessToken: "doesntmatter"},
}

func TestBasicAuthTransport(t *testing.T) {
	bat := &BasicAuthTransport{"user", "pass"}
	client := bat.Client()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Basic dXNlcjpwYXNz" {
			t.Fatal("unexpected auth header")
		}
	}))
	if _, err := client.Get(ts.URL); err != nil {
		t.Fatal(err)
	}
}

func TestQueryDistricts(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	results := clever.QueryAll("/v1.1/districts", nil)

	if !results.Next() {
		t.Fatal("Found no districts")
	}
	district0 := District{}
	if err := results.Scan(&district0); err != nil {
		t.Fatalf("Error retrieving district: %s\n", err)
	}

	resp := DistrictResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/districts/%s", district0.Id), nil, &resp); err != nil {
		t.Fatalf("Error retrieving district: %s\n", err)
	}

	expectedDistrict0 := District{
		Id:   "51a5a56312ec00cc5100007e",
		Name: "test district",
	}
	if !reflect.DeepEqual(expectedDistrict0, district0) {
		t.Fatal("District did not match expected.")
	}
}

func TestQuerySchools(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	results := clever.QueryAll("/v1.1/schools", nil)
	if !results.Next() {
		t.Fatal("Found no schools")
	}
	school0 := School{}
	if err := results.Scan(&school0); err != nil {
		t.Fatalf("Error retrieving school: %s\n", err)
	}

	resp := SchoolResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/schools/%s", school0.Id), nil, &resp); err != nil {
		t.Fatalf("Error retrieving school: %s\n", err)
	}

	expectedSchool0 := School{
		Created:      "2012-11-06T00:00:00Z",
		District:     "51a5a56312ec00cc5100007e",
		HighGrade:    "8",
		Id:           "4fee004cca2e43cf27000002",
		LastModified: "2012-11-07T00:44:53.079Z",
		Location: Location{
			Address: "42139 Fadel Mountains",
			City:    "Thomasfort",
			State:   "AP",
			Zip:     "61397-3760",
		},
		LowGrade:     "6",
		Name:         "Clever Preparatory",
		NcesId:       "94755881",
		Phone:        "(527) 825-2248",
		SchoolNumber: "2559",
		SisId:        "2559",
		StateId:      "23",
	}
	if !reflect.DeepEqual(expectedSchool0, school0) {
		t.Fatalf("School did not match expected.")
	}
}

func TestQueryTeachers(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	results := clever.QueryAll("/v1.1/teachers", nil)
	if !results.Next() {
		t.Fatal("Found no teachers")
	}
	teacher0 := Teacher{}
	if err := results.Scan(&teacher0); err != nil {
		t.Fatalf("Error retrieving teacher: %s\n", err)
	}

	resp := TeacherResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/teachers/%s", teacher0.Id), nil, &resp); err != nil {
		t.Fatalf("Error retrieving teachers: %s\n", err)
	}

	expectedTeacher0 := Teacher{
		Created:      "2013-05-29T06:51:25.139Z",
		District:     "51a5a56312ec00cc5100007e",
		Email:        "t.teacher2@ga4edu.org",
		Id:           "51a5a56d4867bbdf51053c34",
		LastModified: "2013-05-29T06:51:25.145Z",
		Name: Name{
			First:  "Test",
			Middle: "",
			Last:   "Teacher",
		},
		School:        "4fee004cca2e43cf27000002",
		SisId:         "56",
		TeacherNumber: "100",
		Title:         "Math Teacher",
	}
	if !reflect.DeepEqual(expectedTeacher0, teacher0) {
		t.Fatalf("Teacher did not match expected.")
	}
}

func TestQueryStudents(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	results := clever.QueryAll("/v1.1/students", nil)
	if !results.Next() {
		t.Fatal("Found no students")
	}
	student0 := Student{}
	if err := results.Scan(&student0); err != nil {
		t.Fatalf("Error retrieving student: %s\n", err)
	}

	resp := StudentResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/students/%s", student0.Id), nil, &resp); err != nil {
		t.Fatalf("Error retrieving students: %s\n", err)
	}

	expectedStudent0 := Student{
		District:          "51a5a56312ec00cc5100007e",
		Dob:               "12/12/1998",
		FrlStatus:         "Paid",
		Gender:            "M",
		Grade:             "9",
		HispanicEthnicity: "N",
		Race:              "Caucasian",
		School:            "4fee004cca2e43cf27000002",
		SisId:             "1",
		StateId:           "2237504",
		StudentNumber:     "24772",
		Location: Location{
			Address: "",
			City:    "",
			State:   "",
			Zip:     "",
		},
		Name: Name{
			First:  "John",
			Middle: "",
			Last:   "Doe",
		},
		LastModified: "2013-05-29T06:51:28.006Z",
		Created:      "2013-05-29T06:51:27.997Z",
		Email:        "john.doe@ga4edu.org",
		Id:           "51a5a56f4867bbdf51054054",
	}
	if !reflect.DeepEqual(expectedStudent0, student0) {
		t.Fatalf("Student did not match expected.")
	}
}

func TestQuerySections(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	results := clever.QueryAll("/v1.1/sections", nil)
	if !results.Next() {
		t.Fatal("Found no sections")
	}
	section0 := Section{}
	if err := results.Scan(&section0); err != nil {
		t.Fatalf("Error retrieving section: %s\n", err)
	}

	resp := SectionResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/sections/%s", section0.Id), nil, &resp); err != nil {
		t.Fatalf("Error retrieving sections: %s\n", err)
	}

	expectedSection0 := Section{
		CourseName:   "test course",
		CourseNumber: "12345",
		Created:      "2013-05-29T06:51:27.997Z",
		District:     "51a5a56312ec00cc5100007e",
		Grade:        "K",
		Id:           "51a5a56f4867bbdf51054054",
		LastModified: "2013-05-29T06:51:28.006Z",
		Name:         "test section",
		School:       "4fee004cca2e43cf27000002",
		SisId:        "1",
		Students:     []string{"51a5a56f4867bbdf51054054"},
		Subject:      "math",
		Teacher:      "51a5a56d4867bbdf51053c34",
		Term: Term{
			Name:      "",
			StartDate: "",
			EndDate:   "",
		},
	}
	if !reflect.DeepEqual(expectedSection0, section0) {
		t.Fatalf("Section did not match expected.")
	}
}

func postDistrictTest(req *http.Request, params map[string]string) error {
	district := District{}
	if err := json.NewDecoder(req.Body).Decode(&district); err != nil {
		return fmt.Errorf("{\"error\":\"%s\"}", err.Error())
	}
	if district.Name != "new name district" {
		return fmt.Errorf("{\"error\":\"%s\"}", "Error formatting post request")
	}
	return nil
}

func TestPostRequest(t *testing.T) {
	clever := New(mock.NewMock(postDistrictTest, "./data"))
	resp := map[string]string{}
	district1 := District{
		Name: "new name district",
	}
	if err := clever.Request("POST", "/v1.1/districts", nil, district1, &resp); err != nil {
		t.Fatalf("Error posting district: %s\n", err)
	}
}

func TestQueryAll(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	result := clever.QueryAll("/v1.1/sections", nil)

	count := 0
	for result.Next() {
		section := Section{}
		result.Scan(&section)
		count++
	}
	if count != 2 {
		t.Fatalf("Did not get both section pages.")
	}
}

func TestTooManyRequestsError(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	result := clever.QueryAll("/mock/rate/limiter", nil)
	result.Next()
	if result.Error() == nil {
		t.Fatalf("Http response 429 (TooManyRequests) did not trigger an error as expected.")
	} else if _, ok := result.Error().(*TooManyRequestsError); !ok {
		t.Fatalf("Http response 429 (TooManyRequests) did not generate the expected error.")
	}
}

func TestHandlesErrors(t *testing.T) {
	clever := New(mock.NewMock(nil, "./data"))
	result := clever.QueryAll("/mock/error", nil)
	result.Next()
	if result.Error() == nil {
		t.Fatalf("error endpoint did not trigger an error as expected")
	} else if result.Error().Error() != "there was an error (1337)" {
		t.Fatalf("error endpoint did not generate the expected error: %s", result.Error().Error())
	}
}
