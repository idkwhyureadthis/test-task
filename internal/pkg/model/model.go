package model

import "errors"

type Account struct {
	Id      int
	balance float64
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

func NewAccount(id int) Account {
	return Account{Id: id, balance: 0}
}

func (a *Account) GetBalance() float64 {
	return float64(a.balance)
}

func (a *Account) Withdraw(amount float64) error {

	if a.balance < amount {
		return errors.New("not enough money to withdraw")
	}

	if amount <= 0 {
		return errors.New("wrong withdraw sum provided")
	}

	a.balance -= amount

	return nil
}

func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("wrong deposit sum provided")
	}

	a.balance += amount

	return nil
}
