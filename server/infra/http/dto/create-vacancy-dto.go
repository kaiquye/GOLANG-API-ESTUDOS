package dto

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateVacancyDto struct {
	Tilte           string    `json:"tilte" validate:"required"`
	Body            string    `json:"body" validate:"required"`
	Location        int16     `json:"location" validate:"required,number"`
	Level           int16     `json:"level" validate:"required,number"`
	Salary_range    float32   `json:"salary_range" validate:"required,numeric"`
	Organization_id uuid.UUID `json:"organization_id" validate:"required"`
}

func (dto *CreateVacancyDto) Validate(ctx *gin.Context) error {
	err := ctx.BindJSON(&dto)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("sent data is invalid")
	}

	validate := validator.New()
	err = validate.Struct(dto)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
