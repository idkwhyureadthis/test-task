package service

import (
	"errors"
	"strconv"

	"github.com/idkwhyureadthis/test-task/internal/pkg/storage"
)

type Service struct{}

func (s *Service) CreateAccount() int {
	id := storage.CreateNewAccount()
	return id
}

func (s *Service) Deposit(id string, amount string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("wrong id data provided")
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return errors.New("wrong amount data provided")
	}
	return storage.Deposit(idInt, float64(amountFloat))
}

func (s *Service) GetBalance(id string) (float64, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return storage.GetBalance(idInt)
}

func (s *Service) Withdraw(id, amount string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}
	return storage.Withdraw(idInt, amountFloat)
}

func New() *Service {
	return &Service{}
}
