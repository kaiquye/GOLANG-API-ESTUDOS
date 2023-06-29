package mappers

import (
	"fullVagas/domain/models"
	"fullVagas/infra/http/dto"
)

func OrganizationToDomain(data *dto.CreateOrganizationDto) (*models.Organization, error) {
	organization := models.Organization{}
	result, err := organization.OrgPrepare(data.Name, data.Logo, data.Document, &data.Password)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func OrganizationToView(data *models.Organization) models.Organization {
	organization := models.Organization{
		Name:     data.Name,
		Document: data.Document,
		Org_id:   data.Org_id,
		Logo:     data.Logo,
	}
	return organization
}
