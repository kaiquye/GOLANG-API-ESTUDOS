package mappers

import (
	"errors"
	"fullVagas/domain/models"
	"fullVagas/infra/http/dto"
)

func CandidateToDomain(data dto.ApplyToVacancyDto, vacancy models.Vacancy) (*models.Candidate, error) {
	candidate := models.Candidate{}
	result, err := candidate.Prepare(data.Level, data.Location, data.Email, data.Document, data.Resume, vacancy)
	if err != nil {
		return nil, errors.New("error mapping candidate")
	}

	return result, nil
}
