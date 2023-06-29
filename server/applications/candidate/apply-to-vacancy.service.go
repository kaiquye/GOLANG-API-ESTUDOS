package candidate

import (
	IRepository "fullVagas/applications/candidate/repository"
	IRepository2 "fullVagas/applications/vacancy/repository"
	"fullVagas/domain/services/mappers"
	"fullVagas/domain/services/response"
	"fullVagas/infra/database/pg/repository"
	"fullVagas/infra/http/dto"
	"time"
)

type ApplyToVacancyService struct {
	candidateRep IRepository.ICandidateRepository
	vacancyRep   IRepository2.IVacancyRepository
}

func NewApplyToVacancyService() *ApplyToVacancyService {
	return &ApplyToVacancyService{
		candidateRep: repository.NewCandidateRepository(),
		vacancyRep:   repository.NewVacancyRepository(),
	}
}

func (service *ApplyToVacancyService) Run(data dto.ApplyToVacancyDto) response.Template {
	vacancyFound, err := service.vacancyRep.FindById(data.Vacancy_id)
	if err != nil {
		panic(err.Error())
		return response.Template{
			Status:  500,
			Message: "error when searching for a vacancy",
			Date:    time.Now(),
		}
	}
	if vacancyFound == nil {
		return response.Template{
			Status:  404,
			Message: "vacancy not found.",
			Date:    time.Now(),
		}
	}

	mapper, _ := mappers.CandidateToDomain(data, *vacancyFound)
	candidacyApplied, err := service.candidateRep.Apply(*mapper)
	if err != nil {
		return response.Template{
			Status:  500,
			Message: "error when completing the application",
			Date:    time.Now(),
		}
	}

	return response.Template{
		Status: 201,
		Data:   candidacyApplied,
	}
}
