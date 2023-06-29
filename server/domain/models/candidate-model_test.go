package models

import (
	"testing"
)

func TestCandidate_Prepare(t *testing.T) {
	v := &Vacancy{}
	o := &Organization{}
	organization, err := o.OrgPrepare("HFK", "www.logo.com", "00000000000")

	if err != nil {
		t.Errorf("Expected create a organization, but got %s", err.Error())
		t.FailNow()
	}

	vacancy, err := v.Prepare("Developer JS", "vacancy for javascript developer with nodejs and react.", 40000, 4, 2, *organization)

	if err != nil {
		t.Errorf("Expected create a vacancy, but got %s", err.Error())
		t.FailNow()
	}

	c := Candidate{}
	result, err := c.Prepare(1, 2, "kaique@gmail.com", "00000000000", "I'm a javascript developer using nodejs and react", *vacancy)

	expected := &Candidate{
		Level:        1,
		Location:     2,
		Email:        "kaique@gmail.com",
		Document:     "00000000000",
		Resume:       "I'm a javascript developer using nodejs and react",
		Vacancy:      *vacancy,
		candidate_id: result.candidate_id,
		Created_at:   result.Created_at,
		Score:        result.Score,
	}

	if err != nil {
		t.Errorf("Expected %v, but got %s", expected, err.Error())
		t.FailNow()
	}

	if *result != *expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCandidate_CalcScore(t *testing.T) {
	c := Candidate{}
	result, err := c.CalcScore(2, 3)
	expected := 3

	if err != nil {
		t.Errorf("Expected %d, but got %s", expected, err.Error())
		t.FailNow()
	}

	if int(result) != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
