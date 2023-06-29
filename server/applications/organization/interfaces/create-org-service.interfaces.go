package interfaces

import (
	"fullVagas/domain/services/response"
	"fullVagas/infra/http/dto"
)

type ICreateOrganizationService interface {
	Run(data *dto.CreateOrganizationDto) response.Template
}
