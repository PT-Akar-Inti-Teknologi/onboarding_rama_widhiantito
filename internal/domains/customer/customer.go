package customer

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name        string `json:"name" validate:"required"`
	Phonenumber string `json:"phonenumber" validate:"required,len=12"`
	Email       string `json:"email" validate:"required,email"`
	Address     string `json:"address" validate:"required"`
}

func (c *Customer) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
