package vacancy

import (
	IRepository "fullVagas/applications/vacancy/repository"
	"fullVagas/domain/models"
	"fullVagas/domain/services/response"
	"fullVagas/infra/database/pg/repository"
	"time"
)

type ListVacancy struct {
	vacancyRep IRepository.IVacancyRepository
}

func NewListVacancy() *ListVacancy {
	return &ListVacancy{
		vacancyRep: repository.NewVacancyRepository(),
	}
}

func (service *ListVacancy) Run() response.Template {
	vacancys, err := service.vacancyRep.FindAll()
	if err != nil {
		return response.Template{
			Status:  500,
			Message: "internal error",
			Date:    time.Now(),
		}
	}

	if len(vacancys) == 0 {
		return response.Template{
			Status: 200,
			Data:   make([]models.Vacancy, 0),
			Date:   time.Now(),
		}
	}
	return response.Template{
		Status: 200,
		Data:   vacancys,
		Date:   time.Now(),
	}
}
