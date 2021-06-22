package domain

import (
	"github.com/bohdanstryber/banking-go/errs"
	"github.com/bohdanstryber/banking-go/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	if status != "" {
		findAllSQL = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

		err = d.client.Select(&customers, findAllSQL, status)
	} else {
		err = d.client.Select(&customers, findAllSQL)
	}

	if err != nil {
		logger.Error("DB all customers query error" + err.Error())

		return nil, errs.NewUnexpectedError("unexpected db error while processing all customers")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var c Customer

	err := d.client.Get(&c, customerSQL, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("DB scanning customer error" + err.Error())

			return nil, errs.NewUnexpectedError("unexpected db error")
		}

	}

	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}