package service

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/repository"
	"time"
)

type CreateTickerInput struct {
	Symbol string
}

type FetchTickerInput struct {
	DateFrom time.Time
	DateTo   time.Time
	Symbol   string
}

type Tickers interface {
	Create(ctx context.Context, input CreateTickerInput) error
	Fetch(ctx context.Context, input FetchTickerInput) error
}

type Services struct {
	Tickers Tickers
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Tickers: NewTickersService(deps.Repos.Tickers),
	}
}
