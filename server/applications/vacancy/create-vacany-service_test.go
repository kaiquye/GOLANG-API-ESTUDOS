package vacancy

import (
	"fmt"
	repositoryInMemory "fullVagas/infra/database/memory"
	"fullVagas/infra/http/dto"
	"github.com/google/uuid"
	"testing"
)

func TestCreateVacancyService_run(t *testing.T) {
	service := CreateVacancyService{
		vacancyRep:      repositoryInMemory.NewVacancyRepositoryInMemory(),
		organizationRep: repositoryInMemory.NewOrganizationRepositoryInMemory(),
	}
	org_id, err := uuid.Parse("ac3c2675-39ce-4a6b-bff4-f26b74778bce")
	if err != nil {
		panic(err)
	}
	fmt.Println(org_id)
	request := dto.CreateVacancyDto{
		Salary_range:    6000.0,
		Level:           3,
		Body:            "vacancy for administration",
		Location:        4,
		Tilte:           "vacancy for administration",
		Organization_id: org_id,
	}

	result := service.Run(request)
	expected := struct {
		Status int
	}{
		Status: 201,
	}

	if result.Status != expected.Status {
		t.Errorf("Expected %+v, but got %v", expected, result.Status)
	}
}
