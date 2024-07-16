package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/idkwhyureadthis/test-task/internal/endpoint"
	"github.com/idkwhyureadthis/test-task/internal/service"
)

type App struct {
	r *chi.Mux
	e *endpoint.Endpoint
	s *service.Service
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.r = chi.NewRouter()
	a.r.Get("/accounts", a.e.CreateAccount)
	a.r.Post("/accounts/{id}/deposit", a.e.Deposit)
	a.r.Get("/accounts/{id}/balance", a.e.GetBalance)
	a.r.Post("/accounts/{id}/withdraw", a.e.Withdraw)
	return a, nil
}

func (a *App) Run(addr string) error {
	srv := &http.Server{
		Addr:    addr,
		Handler: a.r,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
