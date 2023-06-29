package organization

import (
	"fmt"
	IRepository "fullVagas/applications/organization/repository"
	"fullVagas/domain/services/mappers"
	"fullVagas/domain/services/response"
	"fullVagas/infra/config"
	"fullVagas/infra/database/pg/repository"
	"fullVagas/infra/http/dto"
	"time"
)

type CreateOrganizationService struct {
	organizationRep IRepository.IOrganizationRepository
}

func NewCreateOrganizationService() *CreateOrganizationService {
	return &CreateOrganizationService{
		organizationRep: repository.NewOrganizationRepository(),
	}
}

func (service CreateOrganizationService) Run(data *dto.CreateOrganizationDto) response.Template {
	organizationAlreadyRegistered, err := service.organizationRep.Exist(data.Document)
	if err != nil {
		return response.Template{
			Status:  500,
			Message: err.Error(),
			Date:    time.Now(),
		}
	}
	if organizationAlreadyRegistered != false {
		return response.Template{
			Status:  409,
			Message: "organization already registered.",
			Date:    time.Now(),
		}
	}

	passConfig := config.NewPasswordHash()
	hashedPassword, ok := passConfig.Generate(data.Password)
	if ok == false {
		return response.Template{
			Status:  500,
			Message: "error setting up an organization",
			Date:    time.Now(),
		}
	}

	data.Password = string(*hashedPassword)
	fmt.Println(hashedPassword)
	fmt.Println(data.Password)
	organizationMapper, err := mappers.OrganizationToDomain(data)
	if err != nil {
		return response.Template{
			Status:  400,
			Message: err.Error(),
			Date:    time.Now(),
		}
	}

	newOrganization, err := service.organizationRep.Save(organizationMapper)
	if err != nil {
		return response.Template{
			Status:  500,
			Message: err.Error(),
			Date:    time.Now(),
		}
	}

	return response.Template{
		Status: 201,
		Data:   mappers.OrganizationToView(newOrganization),
		Date:   time.Now(),
	}
}
