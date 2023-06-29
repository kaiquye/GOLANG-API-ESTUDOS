package organization

import (
	IRepository "fullVagas/applications/organization/repository"
	"fullVagas/domain/models"
	"fullVagas/domain/services/response"
	"fullVagas/infra/database/pg/repository"
	"time"
)

type ListAllOrganization struct {
	organizationRep IRepository.IOrganizationRepository
}

func NewListAllOrganization() *ListAllOrganization {
	return &ListAllOrganization{
		organizationRep: repository.NewOrganizationRepository(),
	}
}

func (service *ListAllOrganization) Run() response.Template {
	organizations, err := service.organizationRep.FindAll()
	if err != nil {
		return response.Template{
			Status:  500,
			Message: "internal server error",
			Date:    time.Now(),
		}
	}

	if len(organizations) == 0 {
		return response.Template{
			Status: 200,
			Data:   make([]models.Organization, 0),
			Date:   time.Now(),
		}
	}

	return response.Template{
		Status: 200,
		Data:   organizations,
		Date:   time.Now(),
	}
}
