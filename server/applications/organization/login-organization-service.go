package organization

import (
	IRepository "fullVagas/applications/organization/repository"
	"fullVagas/domain/services/response"
	"fullVagas/infra/config"
	"fullVagas/infra/database/pg/repository"
	"fullVagas/infra/http/dto"
	"time"
)

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type LoginOrganizationService struct {
	organizationRep IRepository.IOrganizationRepository
}

func NewLoginOrganizationService() *LoginOrganizationService {
	return &LoginOrganizationService{
		organizationRep: repository.NewOrganizationRepository(),
	}
}

func (service *LoginOrganizationService) Run(data dto.LoginOrganizationDto) response.Template {
	organizationFound, err := service.organizationRep.FindByDocument(data.Document)
	if err != nil {
		return response.Template{
			Status:  500,
			Message: "error when trying to validate organization",
			Date:    time.Now(),
		}
	}

	if organizationFound == nil {
		return response.Template{
			Status:  404,
			Message: "unauthorized organization",
			Data:    time.Now(),
		}
	}

	passConfig := config.NewPasswordHash()
	match := passConfig.Compare(organizationFound.Password, data.Password)

	if !match {
		return response.Template{
			Status:  404,
			Message: "unauthorized organization",
			Data:    time.Now(),
		}
	}

	tokenConfig := config.NewAccessTokenConfig()
	claims := tokenConfig.FormatClaims(map[string]interface{}{
		"document": organizationFound.Document,
		"org_id":   organizationFound.Org_id,
	})
	token := tokenConfig.Generate(claims)

	result := AccessTokenResponse{
		AccessToken: token,
	}
	return response.Template{
		Status: 200,
		Data:   result,
	}
}
