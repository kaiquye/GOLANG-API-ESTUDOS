package IRepository

import (
	"fullVagas/domain/models"
	"github.com/google/uuid"
)

type IVacancyRepository interface {
	Create(data *models.Vacancy) (*models.Vacancy, error)
	FindById(id uuid.UUID) (*models.Vacancy, error)
	FindAll() ([]models.Vacancy, error)
}
