package dto

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateOrganizationDto struct {
	Name     string `json:"name" validate:"required"`
	Logo     string `json:"logo" validate:"required"`
	Document string `json:"document" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto *CreateOrganizationDto) Validate(ctx *gin.Context) error {
	err := ctx.BindJSON(&dto)
	if err != nil {
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
