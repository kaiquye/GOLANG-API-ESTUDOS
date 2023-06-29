package interfaces

import (
	"fullVagas/domain/services/response"
	"fullVagas/infra/http/dto"
)

type ICreateVacancyService interface {
	Run(data *dto.CreateVacancyDto) response.Template
}
