package dto

import (
	"github.com/bohdanstryber/banking-go/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type NewTransactionRequest struct {
	AccountId string `json:"account_id"`
	CustomerId string `json:"customer_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can be DEPOSIT or WITHDRAWAL")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount can't be less than zero")
	}

	return nil
}

func (r NewTransactionRequest) IsWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}

	return false
}