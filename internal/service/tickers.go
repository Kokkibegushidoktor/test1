package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"github.com/Kokkibegushidoktor/test1/internal/repository"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const _workerInterval = time.Minute * 1

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

func (s *TickersService) GetAllTickers(ctx context.Context) ([]models.Ticker, error) {
	tickers, err := s.repo.GetAllTickers(ctx)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}

func (s *TickersService) Fetch(ctx context.Context, inp FetchTickerInput) (models.FetchResponse, error) {
	rates, err := s.repo.FetchFromTo(ctx, inp.Symbol, inp.DateFrom, inp.DateTo)
	if err != nil {
		return models.FetchResponse{}, err
	}
	if len(rates) < 1 {
		return models.FetchResponse{}, models.ErrRatesNotFound
	}

	priceFrom, err := strconv.ParseFloat(rates[0].Price, 64)
	priceTo, err := strconv.ParseFloat(rates[len(rates)-1].Price, 64)

	dif := priceTo - priceFrom
	dif = dif * 100 / priceFrom

	return models.FetchResponse{
		Ticker:     inp.Symbol,
		Price:      priceTo,
		Difference: dif,
	}, nil
}

func (s *TickersService) InitUpdateWorker(ctx context.Context) {
	go s.processUpdateRates(ctx)
}

func (s *TickersService) processUpdateRates(ctx context.Context) {
	for {
		s.updateRates(ctx)
		time.Sleep(_workerInterval)
	}
}

func (s *TickersService) updateRates(ctx context.Context) {
	tickers, err := s.GetAllTickers(ctx)
	if err != nil {
		if !errors.Is(err, models.ErrTickerNotFound) {
			log.Error().Msgf("error getting all tickers, err: %v", err)
		}
		return
	}

	baseEndpoint := "https://api1.binance.com"
	var symbols strings.Builder //"[\"BTCUSDT\",\"BNBUSDT\"]"
	symbols.WriteString("[")
	for i, ticker := range tickers {
		if i == len(tickers)-1 {
			symbols.WriteString(fmt.Sprintf("\"%s\"", ticker.Symbol))
		} else {
			symbols.WriteString(fmt.Sprintf("\"%s\",", ticker.Symbol))
		}
	}
	symbols.WriteString("]")
	res, err := http.Get(fmt.Sprintf("%s/api/v3/ticker/price?symbols=%s", baseEndpoint, symbols.String()))
	if err != nil {
		log.Error().Msgf("error getting price information, err: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Msgf("error reading price information, err: %v", err)
	}

	rates := make([]*models.Rate, 1)
	if err = json.Unmarshal(body, &rates); err != nil {
		log.Error().Msgf("error unmarshalling price information, err: %v", err)
	}

	if err != nil {
		return
	}

	for _, rate := range rates {
		rate.Time = time.Now()
	}

	if err = s.repo.AddRates(ctx, rates); err != nil {
		log.Error().Msgf("error writing price information, err: %v", err)
	}

}
