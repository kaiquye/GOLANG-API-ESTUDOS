package repositoryInMemory

import (
	"fullVagas/domain/models"
	"github.com/google/uuid"
)

type OrganizationRepositoryInMemory struct {
	databaseInMemory []models.Organization
}

func NewOrganizationRepositoryInMemory() *OrganizationRepositoryInMemory {
	return &OrganizationRepositoryInMemory{
		databaseInMemory: make([]models.Organization, 0),
	}
}

func (rep *OrganizationRepositoryInMemory) Save(data *models.Organization) (*models.Organization, error) {
	rep.databaseInMemory = append(rep.databaseInMemory, *data)
	return data, nil
}

func (rep *OrganizationRepositoryInMemory) Exist(document string) (bool, error) {
	for _, organization := range rep.databaseInMemory {
		if organization.Document == document {
			return true, nil
		}
	}
	return false, nil
}

func (rep *OrganizationRepositoryInMemory) FindById(id uuid.UUID) (bool, error) {
	for _, organization := range rep.databaseInMemory {
		if organization.Org_id == id {
			return true, nil
		}
	}
	return false, nil
}
