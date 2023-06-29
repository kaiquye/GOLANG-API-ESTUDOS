package interfaces

import (
	"fullVagas/domain/services/response"
	"fullVagas/infra/http/dto"
)

type IApplyToVacancyService interface {
	Run(data dto.ApplyToVacancyDto) response.Template
}
