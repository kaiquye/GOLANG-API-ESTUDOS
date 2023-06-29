package IRepository

import (
	"fullVagas/domain/models"
	"github.com/google/uuid"
)

type IOrganizationRepository interface {
	Save(data *models.Organization) (*models.Organization, error)
	Exist(document string) (bool, error)
	FindById(id uuid.UUID) (bool, error)
	FindByDocument(document string) (*models.Organization, error)
	FindAll() ([]models.Organization, error)
}
