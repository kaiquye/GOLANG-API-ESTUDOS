package models

import (
	"github.com/google/uuid"
	"time"
)

type Vacancy struct {
	Vacancy_id   uuid.UUID
	Organization Organization `json:"organization" validate:"required"`
	Tilte        string       `json:"tilte" validate:"required"`
	Body         string       `json:"body" validate:"required"`
	Location     int16        `json:"location" validate:"required"`
	Level        int16        `json:"level" validate:"required"`
	Salary_range float32      `json:"salary_range" validate:"required"`
	Created_at   time.Time    `json:"created_at" validate:"required"`
}

func (v *Vacancy) Prepare(tilte, body string, salary_range float32, level, location int16, organization Organization) (*Vacancy, error) {
	vacany := &Vacancy{
		Vacancy_id:   uuid.New(),
		Body:         body,
		Tilte:        tilte,
		Organization: organization,
		Salary_range: salary_range,
		Location:     location,
		Level:        level,
	}

	return vacany, nil
}
