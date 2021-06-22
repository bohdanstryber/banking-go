package domain

import (
	"github.com/bohdanstryber/banking-go/dto"
	"github.com/bohdanstryber/banking-go/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(id string) (*Account, *errs.AppError)
	NewTransaction(t Transaction) (*Transaction, *errs.AppError)
}