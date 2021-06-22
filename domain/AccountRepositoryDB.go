package domain

import (
	"github.com/bohdanstryber/banking-go/errs"
	"github.com/bohdanstryber/banking-go/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("error while creating new account:" + err.Error())

		return nil, errs.NewUnexpectedError("error while creating new account:" + err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id:" + err.Error())

		return nil, errs.NewUnexpectedError("error while getting last insert id:" + err.Error())
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDB) FindById(id string) (*Account, *errs.AppError) {
	sql := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"
	var a Account
	err := d.client.Get(&a, sql, id)

	if err != nil {
		logger.Error("DB error while fetching transaction info:" + err.Error())

		return nil, errs.NewUnexpectedError("unexpected db error")
	}

	return &a, nil
}

func (d AccountRepositoryDB) NewTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting new transaction for bank account transaction: " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
	res, _ := tx.Exec(sqlInsert,
		t.AccountId,
		t.Amount,
		t.TransactionType,
		t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? where account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? where account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new transaction:" + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected db err (Error while creating new transaction):" + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting new transaction:" + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected db err: 	" + err.Error())

	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("EError while getting last insert id (transaction):" + err.Error())

		return nil, errs.NewUnexpectedError("Error while getting last insert id (transaction):" + err.Error())
	}

	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(id, 10)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}