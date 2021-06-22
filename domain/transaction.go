package domain

import "github.com/bohdanstryber/banking-go/dto"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type Transaction struct {
	TransactionId string `db:"transaction_id"`
	AccountId string `db:"account_id"`
	Amount float64 `db:"amount"`
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}

	return false
}

func (t Transaction) ToDtoResponse() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}