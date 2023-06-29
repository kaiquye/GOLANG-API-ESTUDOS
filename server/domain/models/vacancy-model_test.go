package models

import "testing"

func TestVacancy_Prepare(t *testing.T) {
	vacancy := Vacancy{}
	org := Organization{}
	organization, err := org.OrgPrepare("HFK", "www.logo.com", "00000000000")

	if err != nil {
		t.Errorf("Expected create a organization, but got %s", err.Error())
		t.FailNow()
	}

	result, err := vacancy.Prepare("Developer JS", "vacancy for javascript developer with nodejs and react.", 40000, 4, 2, *organization)

	if err != nil {
		t.Errorf("Expected create a user, but got %s", err.Error())
		t.FailNow()
	}

	expected := &Vacancy{
		Level:        4,
		Location:     2,
		Tilte:        "Developer JS",
		Body:         "vacancy for javascript developer with nodejs and react.",
		Salary_range: 40000,
		Organization: *organization,
		Vacancy_id:   result.Vacancy_id,
	}

	if *result != *expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
