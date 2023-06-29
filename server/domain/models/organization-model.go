package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Organization struct {
	Org_id   uuid.UUID
	Name     string `json:"name" validate:"required"`
	Logo     string `json:"logo"`
	Document string `json:"document" validate:"required"`
	Password string `json:"password"`
}

func (org *Organization) OrgPrepare(name, logo, document string, password *string) (*Organization, error) {
	fmt.Println(password)
	OrgPrepare := &Organization{
		Org_id:   uuid.New(),
		Name:     name,
		Logo:     logo,
		Document: document,
		Password: *password,
	}

	_, err := OrgPrepare.ValidateDocument()
	if err != nil {
		return nil, err
	}

	return OrgPrepare, nil
}

func (org *Organization) ValidateDocument() (ok bool, typeError error) {
	if len(org.Document) < 11 {
		return false, errors.New("invalid document")
	}
	return true, nil
}
