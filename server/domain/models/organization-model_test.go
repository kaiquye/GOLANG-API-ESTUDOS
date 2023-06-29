package models

import (
	"errors"
	"testing"
)

func TestOrganization_OrgPrepare(t *testing.T) {
	org := Organization{}
	result, err := org.OrgPrepare("kaique", "/url.image.com", "00000000000")
	expected := &Organization{
		Name:     "kaique",
		Logo:     "/url.image.com",
		Document: "00000000000",
		Org_id:   result.Org_id,
	}

	if err != nil {
		t.Errorf("")
	}

	if *result != *expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

}

func TestOrganization_ValidateDocument(t *testing.T) {
	org := Organization{}
	_, err := org.OrgPrepare("kaique", "/url.image.com", "0000000000")
	expected := errors.New("invalid document")

	if err.Error() != expected.Error() {
		t.Errorf("Expected %s, but got %s", expected, err)
	}
}
