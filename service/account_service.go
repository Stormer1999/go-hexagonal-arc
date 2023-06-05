package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// validate input
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}

	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type not valid")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := accountRequestDTOWrapper(newAcc)

	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		accDTO := accountRequestDTOWrapper(&account)
		responses = append(responses, accDTO)
	}

	return responses, nil
}

func accountRequestDTOWrapper(acc *repository.Account) AccountResponse {
	res := AccountResponse{
		AccountID:   acc.AccountID,
		OpeningDate: acc.OpeningDate,
		AccountType: acc.AccountType,
		Amount:      acc.Amount,
		Status:      acc.Status,
	}

	return res
}
