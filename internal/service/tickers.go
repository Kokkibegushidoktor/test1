package service

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"github.com/Kokkibegushidoktor/test1/internal/repository"
)

type TickersService struct {
	repo repository.Tickers
}

func NewTickersService(repo repository.Tickers) *TickersService {
	return &TickersService{
		repo: repo,
	}
}

func (s *TickersService) Create(ctx context.Context, inp CreateTickerInput) error {
	ticker := models.Ticker{
		Symbol: inp.Symbol,
	}

	return s.repo.Create(ctx, &ticker)
}

func (s *TickersService) Fetch(ctx context.Context, inp FetchTickerInput) (models.FetchResponse, error) {
	rates, err := s.repo.FetchFromTo(ctx, inp.Symbol, inp.DateFrom, inp.DateTo)
	if err != nil {
		return models.FetchResponse{}, err
	}

	priceFrom := rates[0].Price
	priceTo := rates[len(rates)-1].Price

	dif := priceFrom - priceTo
	dif = dif * 100 / priceFrom

	return models.FetchResponse{
		Ticker:     inp.Symbol,
		Price:      priceTo,
		Difference: dif,
	}, nil
}
