package dto

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ApplyToVacancyDto struct {
	Vacancy_id uuid.UUID `json:"vacancy_id" validate:"required"`
	Document   string    `json:"document" validate:"required"`
	Email      string    `json:"email" validate:"required"`
	Level      int16     `json:"level" validate:"required"`
	Location   int16     `json:"location" validate:"required"`
	Resume     string    `json:"resume" validate:"required"`
}

func (dto *ApplyToVacancyDto) Validate(ctx *gin.Context) error {
	err := ctx.BindJSON(&dto)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("send data is invalid")
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		fmt.Println(err)
		return errors.New(err.Error())
	}

	return nil
}
