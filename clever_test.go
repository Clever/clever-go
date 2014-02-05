package clever

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueryDistricts(t *testing.T) {
	clever := NewMock("./data")
	resp1 := &DistrictsResp{}
	if err := clever.Query("/v1.1/districts", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving districts: %s\n", err))
	}
	resp2 := &DistrictResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/districts/%s", resp1.Districts[0].District.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving district: %s\n", err))
	}

	District0 := District{
		Id:   "51a5a56312ec00cc5100007e",
		Name: "test district",
	}
	if !reflect.DeepEqual(District0, resp1.Districts[0].District) {
		t.Error(fmt.Errorf("District did not match expected."))
	}
}

func TestQuerySchools(t *testing.T) {
	clever := NewMock("./data")
	resp1 := &SchoolsResp{}
	if err := clever.Query("/v1.1/schools", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving schools: %s\n", err))
	}
	resp2 := &SchoolResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/schools/%s", resp1.Schools[0].School.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving school: %s\n", err))
	}

	School0 := School{
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
	if !reflect.DeepEqual(School0, resp1.Schools[0].School) {
		t.Error(fmt.Errorf("School did not match expected."))
	}
}

func TestQueryTeachers(t *testing.T) {
	clever := NewMock("./data")
	resp1 := &TeachersResp{}
	if err := clever.Query("/v1.1/teachers", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving teachers: %s\n", err))
	}
	resp2 := &TeacherResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/teachers/%s", resp1.Teachers[0].Teacher.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving teacher: %s\n", err))
	}

	Teacher0 := Teacher{
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
	if !reflect.DeepEqual(Teacher0, resp1.Teachers[0].Teacher) {
		t.Error(fmt.Errorf("Student did not match expected."))
	}
}

func TestQueryStudents(t *testing.T) {
	clever := NewMock("./data")
	resp1 := &StudentsResp{}
	if err := clever.Query("/v1.1/students", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving students: %s\n", err))
	}
	resp2 := &StudentResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/students/%s", resp1.Students[0].Student.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving student: %s\n", err))
	}

	Student0 := Student{
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
	if !reflect.DeepEqual(Student0, resp1.Students[0].Student) {
		t.Error(fmt.Errorf("Student did not match expected."))
	}
}

func TestQuerySections(t *testing.T) {
	clever := NewMock("./data")
	resp1 := &SectionsResp{}
	if err := clever.Query("/v1.1/sections", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving sections: %s\n", err))
	}
	resp2 := &SectionResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/sections/%s", resp1.Sections[0].Section.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving section: %s\n", err))
	}

	Section0 := Section{
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
	if !reflect.DeepEqual(Section0, resp1.Sections[0].Section) {
		t.Error(fmt.Errorf("Student did not match expected."))
	}
}

func TestQueryAll(t *testing.T) {
	clever := NewMock("./data")
	result := clever.QueryAll("/v1.1/sections", nil)

	count := 0
	for result.Next() {
		section := &Section{}
		result.Scan(section)
		count++
	}

	if count != 2 {
		t.Error(fmt.Errorf("Did not get both section pages."))
	}
}
