package models

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Candidate struct {
	Candidate_id uuid.UUID `json:"candidate_id" validate:"required"`
	Vacancy      Vacancy   `json:"vacancy"`
	Document     string    `json:"document" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Level        int16     `json:"level" validate:"required"`
	Location     int16     `json:"location" validate:"required"`
	Resume       string    `json:"resume" validate:"required"`
	Score        int16     `json:"score" validate:"required"`
	Created_at   time.Time `json:"created_at" validate:"required"`
}

func (c *Candidate) Prepare(level, location int16, email, document, resume string, vacancy Vacancy) (candidate *Candidate, error error) {
	score, err := c.CalcScore(level, location)
	if err != nil {
		return nil, err
	}

	CandidatePrepare := &Candidate{
		Candidate_id: uuid.New(),
		Vacancy:      vacancy,
		Level:        level,
		Location:     location,
		Score:        score,
		Resume:       resume,
		Email:        email,
		Document:     document,
		Created_at:   time.Now(),
	}

	return CandidatePrepare, nil
}

func (c *Candidate) CalcScore(level int16, location int16) (score int16, error error) {
	if level < 0 || location < 0 {
		return 0, errors.New("unable to calculate the score, invalid input values")
	}

	result := (level * location) / level
	return result, nil
}
