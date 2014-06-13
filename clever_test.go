package clever

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmptyAuthError(t *testing.T) {
	clever := NewMock("./data")
	clever.Auth = Auth{}
	results := clever.QueryAll("/v1.1/districts", nil)
	if results.Next(); results.lastError == nil {
		t.Fatalf("Empty auth credentials raised no error")
	} else if fmt.Sprint(results.lastError) != "Must provide either API key or bearer token" {
		t.Fatalf("Empty auth credentials raised incorrect error '" + fmt.Sprint(results.lastError) + "'")
	}
}

// helper function for TestBearerAuth and TestBasicAuth
func verifyAuthHeader(auth Auth, expectedAuthString string, t *testing.T) {
	reqHeader := new(map[string][]string)
	clever := NewMock("./data", reqHeader)
	clever.Auth = auth
	results := clever.QueryAll("/v1.1/districts", nil)
	results.Next()
	if header, ok := (*reqHeader)["Authorization"]; ok {
		if len(header) == 0 || header[0] != expectedAuthString {
			t.Fatalf("Expected auth header '" + expectedAuthString + "'; found '" + header[0] + "'")
		}
	} else {
		t.Fatalf("Auth request header not present")
	}
}

func TestBearerAuth(t *testing.T) {
	verifyAuthHeader(Auth{"", "TEST_TOKEN"}, "Bearer TEST_TOKEN", t)
}

func TestBasicAuth(t *testing.T) {
	verifyAuthHeader(Auth{"TEST_KEY", ""}, "Basic VEVTVF9LRVk6", t)
}

func TestQueryDistricts(t *testing.T) {
	clever := NewMock("./data")
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
	clever := NewMock("./data")
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
	clever := NewMock("./data")
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
	clever := NewMock("./data")
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
	clever := NewMock("./data")
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

func TestQueryAll(t *testing.T) {
	clever := NewMock("./data")
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
