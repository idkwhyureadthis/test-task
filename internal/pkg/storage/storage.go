package storage

import (
	"errors"
	"sync"

	"github.com/idkwhyureadthis/test-task/internal/pkg/model"
)

var lastCreatedInd int
var users map[int]*model.Account
var mu sync.Mutex

func Init() {
	users = make(map[int]*model.Account)
	lastCreatedInd = 1
}

func CreateNewAccount() int {
	mu.Lock()
	defer mu.Unlock()
	newAcc := model.NewAccount(lastCreatedInd)
	users[lastCreatedInd] = &newAcc
	lastCreatedInd++
	return newAcc.Id
}

func Deposit(id int, amount float64) error {
	err := make(chan error)
	defer close(err)
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if account, ok := users[id]; ok {
			err <- account.Deposit(amount)
		} else {
			err <- errors.New("user with such id not found")
		}
	}()
	return <-err
}

func Withdraw(id int, amount float64) error {
	err := make(chan error)
	defer close(err)
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if account, ok := users[id]; ok {
			err <- account.Withdraw(amount)
		} else {
			err <- errors.New("user with such id not found")
		}
	}()
	return <-err
}

func GetBalance(id int) (float64, error) {
	type result struct {
		Error error
		Value float64
	}
	res := make(chan result)
	defer close(res)
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if account, ok := users[id]; ok {
			res <- result{Error: nil, Value: account.GetBalance()}
			return
		} else {
			res <- result{Error: errors.New("user not found"), Value: 0}
		}
	}()
	r := <-res
	return r.Value, r.Error
}
