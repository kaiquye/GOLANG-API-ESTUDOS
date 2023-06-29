package dto

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginOrganizationDto struct {
	Document string `json:"document" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto *LoginOrganizationDto) Validate(ctx *gin.Context) error {
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
