package endpoint

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkwhyureadthis/test-task/internal/pkg/formatter"
)

type Service interface {
	CreateAccount() int
	Deposit(id, amount string) error
	GetBalance(id string) (float64, error)
	Withdraw(id, amount string) error
}

type Endpoint struct {
	s Service
}

func (e *Endpoint) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id := e.s.CreateAccount()
	type response struct {
		Id int `json:"user_id"`
	}
	resp := response{Id: id}
	dat, err := json.Marshal(resp)
	if err != nil {
		formatter.JsonifyError(w, err)
		log.Print("failed to marshal json", err)
		return
	}
	log.Printf("POST CREATE_USER %d SUCCESS", id)
	w.WriteHeader(200)
	w.Write(dat)
}

func (e *Endpoint) Deposit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	amount := r.URL.Query().Get("amount")
	err := e.s.Deposit(id, amount)
	if err != nil {
		formatter.JsonifyError(w, err)
		log.Printf("POST user_id: %s DEPOSIT FAILED", id)
		return
	}
	formatter.JsonifyMessage(w, "deposit succesful")
	log.Printf("POST user_id: %s DEPOSIT SUCCESS", id)
}

func (e *Endpoint) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	balance, err := e.s.GetBalance(id)
	if err != nil {
		formatter.JsonifyError(w, err)
		log.Printf("GET user_id: %d GET_BALANCE FAILED", err)
		return
	}
	type response struct {
		Id      int     `json:"user_id"`
		Balance float64 `json:"balance"`
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		formatter.JsonifyError(w, err)
		return
	}
	resp := response{Id: idInt, Balance: balance}
	dat, err := json.Marshal(resp)
	if err != nil {
		formatter.JsonifyError(w, err)
		log.Printf("GET user_id: %d GET_BALANCE FAILED", idInt)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
	log.Printf("GET user_id: %d GET_BALANCE SUCCESS", idInt)
}

func (e *Endpoint) Withdraw(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	amount := r.URL.Query().Get("amount")
	err := e.s.Withdraw(id, amount)
	if err != nil {
		formatter.JsonifyError(w, err)
		log.Printf("POST user_id: %s WITHDRAW FAILED", id)
		return
	}
	log.Printf("POST user_id: %s WITHDRAW SUCCESS", id)
	formatter.JsonifyMessage(w, "WITHDRAW SUCCESS")
}

func New(s Service) *Endpoint {
	return &Endpoint{s: s}
}
