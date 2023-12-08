package customer

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
}
