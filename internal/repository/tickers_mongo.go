package repository

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type TickerRepo struct {
	tickerCollection *mongo.Collection
	ratesCollection  *mongo.Collection
}

func NewTickerRepo(db *mongo.Database) *TickerRepo {
	tickerIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "symbol", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection("tickers").Indexes().CreateOne(context.Background(), tickerIndex)
	if err != nil {
		log.Fatal().Msgf("error creating ticker collection index, err: %v", err)
	}

	return &TickerRepo{
		tickerCollection: db.Collection("tickers"),
		ratesCollection:  db.Collection("rates"),
	}
}
func (r *TickerRepo) Create(ctx context.Context, ticker *models.Ticker) error {
	_, err := r.tickerCollection.InsertOne(ctx, ticker)
	if mongo.IsDuplicateKeyError(err) {
		return models.ErrTickerAlreadyExists
	}

	return err
}

func (r *TickerRepo) AddRate(ctx context.Context, rate *models.Rate) error {
	_, err := r.tickerCollection.InsertOne(ctx, rate)

	return err
}

func (r *TickerRepo) GetBySymbol(ctx context.Context, symbol string) (*models.Ticker, error) {
	var ticker models.Ticker
	if err := r.tickerCollection.FindOne(ctx, bson.M{"symbol": symbol}).Decode(&ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}

func (r *TickerRepo) FetchFromTo(ctx context.Context, symbol string, from, to time.Time) ([]models.Rate, error) {
	query := bson.M{
		"$and": bson.A{
			bson.M{"timestamp": bson.M{"$gte": from}},
			bson.M{"timestamp": bson.M{"$lte": to}},
		},
	}

	cur, err := r.ratesCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	var rates []models.Rate
	if err = cur.All(ctx, &rates); err != nil {
		return nil, err
	}

	return rates, nil
}
