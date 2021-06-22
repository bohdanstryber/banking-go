package service

import (
	"github.com/bohdanstryber/banking-go/domain"
	"github.com/bohdanstryber/banking-go/dto"
	"github.com/bohdanstryber/banking-go/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()

	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil	
}

func (s DefaultAccountService) NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := request.Validate()

	if err != nil {
		return nil, err
	}

	if request.IsWithdrawal() {
		account, err := s.repo.FindById(request.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(request.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.repo.NewTransaction(t)

	if err != nil {
		return nil, err
	}

	response := newTransaction.ToDtoResponse()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}