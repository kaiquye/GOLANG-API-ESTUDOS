package repositoryInMemory

import "fullVagas/domain/models"

type VacancyRepositoryInMemory struct {
	databaseInMemory []models.Vacancy
}

func NewVacancyRepositoryInMemory() *VacancyRepositoryInMemory {
	return &VacancyRepositoryInMemory{
		databaseInMemory: make([]models.Vacancy, 0),
	}
}

func (rep *VacancyRepositoryInMemory) Create(data *models.Vacancy) (*models.Vacancy, error) {
	rep.databaseInMemory = append(rep.databaseInMemory, *data)
	return data, nil
}
