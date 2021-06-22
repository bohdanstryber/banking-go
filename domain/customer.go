package domain

import (
	"github.com/bohdanstryber/banking-go/dto"
	"github.com/bohdanstryber/banking-go/errs"
)

type Customer struct {
	Id string `db:"customer_id"`
	Name string `json:"full_name"`
	City string `json:"city"`
	ZipCode string `json:"zip_code"`
	DateOfBirth string `db:"date_of_birth"`
	Status string
}

func (c Customer) statusAsText() string {
	statusAsText := "inactive"
	if c.Status == "1" {
		statusAsText = "active"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}