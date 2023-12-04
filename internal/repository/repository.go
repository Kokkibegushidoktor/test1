package repository

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Tickers interface {
	Create(ctx context.Context, ticker *models.Ticker) error
	AddRate(ctx context.Context, rate *models.Rate) error
	GetBySymbol(ctx context.Context, symbol string) (*models.Ticker, error)
	FetchFromTo(ctx context.Context, symbol string, from, to time.Time) ([]models.Rate, error)
}

type Repositories struct {
	Tickers Tickers
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Tickers: NewTickerRepo(db),
	}
}
