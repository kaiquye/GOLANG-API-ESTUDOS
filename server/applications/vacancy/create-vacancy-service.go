package vacancy

import (
	"fmt"
	IRepository_org "fullVagas/applications/organization/repository"
	IRepository "fullVagas/applications/vacancy/repository"
	"fullVagas/domain/services/mappers"
	"fullVagas/domain/services/response"
	"fullVagas/infra/database/pg/repository"
	"fullVagas/infra/http/dto"
	"time"
)

type CreateVacancyService struct {
	vacancyRep      IRepository.IVacancyRepository
	organizationRep IRepository_org.IOrganizationRepository
}

func NewCreateVacancyService() *CreateVacancyService {
	return &CreateVacancyService{
		vacancyRep:      repository.NewVacancyRepository(),
		organizationRep: repository.NewOrganizationRepository(),
	}
}

func (service CreateVacancyService) Run(data *dto.CreateVacancyDto) response.Template {
	fmt.Println(data.Organization_id)
	organizationFound, _ := service.organizationRep.FindById(data.Organization_id)
	if organizationFound == false {
		return response.Template{
			Status:  404,
			Message: "organization not found",
			Date:    time.Now(),
		}
	}

	VacancyMapper, err := mappers.VacancyToDomain(data)
	if err != nil {
		return response.Template{
			Status:  400,
			Message: err.Error(),
			Date:    time.Now(),
		}
	}

	NewVacancy, err := service.vacancyRep.Create(VacancyMapper)
	if err != nil {
		return response.Template{
			Status:  500,
			Message: err.Error(),
			Date:    time.Now(),
		}
	}

	return response.Template{
		Status: 201,
		Data:   NewVacancy,
		Date:   time.Now(),
	}
}
