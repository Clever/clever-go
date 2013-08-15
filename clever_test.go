package clever

import (
	"fmt"
	"testing"
)

func TestQueryDistricts(t *testing.T) {
	clever := New(Auth{"DEMO_KEY", ""}, "https://api.getclever.com")
	resp1 := &DistrictsResp{}
	if err := clever.Query("/v1.1/districts", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving districts: %s\n", err))
	}
	resp2 := &DistrictResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/districts/%s", resp1.Districts[0].District.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving district: %s\n", err))
	}
}

func TestQuerySchools(t *testing.T) {
	clever := New(Auth{"DEMO_KEY", ""}, "https://api.getclever.com")
	resp1 := &SchoolsResp{}
	if err := clever.Query("/v1.1/schools", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving schools: %s\n", err))
	}
	resp2 := &SchoolResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/schools/%s", resp1.Schools[0].School.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving school: %s\n", err))
	}
}

func TestQueryTeachers(t *testing.T) {
	clever := New(Auth{"DEMO_KEY", ""}, "https://api.getclever.com")
	resp1 := &TeachersResp{}
	if err := clever.Query("/v1.1/teachers", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving teachers: %s\n", err))
	}
	resp2 := &TeacherResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/teachers/%s", resp1.Teachers[0].Teacher.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving teacher: %s\n", err))
	}
}

func TestQueryStudents(t *testing.T) {
	clever := New(Auth{"DEMO_KEY", ""}, "https://api.getclever.com")
	resp1 := &StudentsResp{}
	if err := clever.Query("/v1.1/students", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving students: %s\n", err))
	}
	resp2 := &StudentResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/students/%s", resp1.Students[0].Student.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving student: %s\n", err))
	}
}

func TestQuerySections(t *testing.T) {
	clever := New(Auth{"DEMO_KEY", ""}, "https://api.getclever.com")
	resp1 := &SectionsResp{}
	if err := clever.Query("/v1.1/sections", map[string]string{}, resp1); err != nil {
		t.Error(fmt.Errorf("Error retrieving sections: %s\n", err))
	}
	resp2 := &SectionResp{}
	if err := clever.Query(fmt.Sprintf("/v1.1/sections/%s", resp1.Sections[0].Section.Id), map[string]string{}, resp2); err != nil {
		t.Error(fmt.Errorf("Error retrieving section: %s\n", err))
	}
}
