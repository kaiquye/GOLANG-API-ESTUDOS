package mappers

import (
	"fullVagas/domain/models"
	"fullVagas/infra/http/dto"
)

func VacancyToDomain(data *dto.CreateVacancyDto) (*models.Vacancy, error) {
	vacancy := models.Vacancy{}
	result, err := vacancy.Prepare(data.Tilte, data.Body, data.Salary_range, data.Level, data.Location, models.Organization{Org_id: data.Organization_id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
