package organization

import (
	"fmt"
	"fullVagas/domain/services/response"
	"fullVagas/infra/database/memory"
	"fullVagas/infra/http/dto"
	"testing"
)

func TestCreateOrganizationService_Run(t *testing.T) {
	fmt.Println("Organization creation test. Expecting to return status 201")
	service := CreateOrganizationService{
		organizationRep: repositoryInMemory.NewOrganizationRepositoryInMemory(),
	}

	request := dto.CreateOrganizationDto{
		Name:     "Kaic",
		Document: "00000000000",
		Logo:     "www.logo.com",
	}

	response := service.Run(&request)
	expected := 201

	if response.Status != expected {
		t.Errorf("Expected %v, but got %v", expected, response.Status)
	}
}

func TestCreateOrganizationService_Run_AlreadyExist(t *testing.T) {
	fmt.Println("Organization creation test. Expecting to return organization already exists")
	service := CreateOrganizationService{
		organizationRep: repositoryInMemory.NewOrganizationRepositoryInMemory(),
	}

	request := dto.CreateOrganizationDto{
		Name:     "Kaic",
		Document: "00000000000",
		Logo:     "www.logo.com",
	}

	service.Run(&request)
	result := service.Run(&request)
	expected := response.Template{
		Status:  409,
		Message: "organization already registered.",
		Data:    nil,
	}

	if result.Status != expected.Status {
		t.Errorf("Expected %v, but got %v", expected.Status, result.Status)
	}

	if result.Message != expected.Message {
		t.Errorf("Expected %v, but got %v", expected.Message, result.Message)
	}
}

func TestCreateOrganizationService_Run_Status_400(t *testing.T) {
	fmt.Println("Organization creation test. Expecting to return status 400")
	service := CreateOrganizationService{
		organizationRep: repositoryInMemory.NewOrganizationRepositoryInMemory(),
	}

	request := dto.CreateOrganizationDto{
		Name:     "Kaic",
		Document: "1234",
		Logo:     "www.logo.com",
	}

	result := service.Run(&request)
	expected := response.Template{
		Status:  400,
		Message: "invalid document",
		Data:    nil,
	}

	if result.Status != expected.Status {
		t.Errorf("Expected %v, but got %v", expected.Status, result.Status)
	}

	if result.Message != expected.Message {
		t.Errorf("Expected %v, but got %v", expected.Message, result.Message)
	}
}
