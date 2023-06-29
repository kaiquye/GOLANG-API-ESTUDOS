package interfaces

import (
	"fullVagas/domain/services/response"
	"fullVagas/infra/http/dto"
)

type ILoginOrganizationService interface {
	Run(data dto.LoginOrganizationDto) response.Template
}
