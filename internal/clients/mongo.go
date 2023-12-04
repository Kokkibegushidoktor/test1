package clients

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/config"
	"github.com/Kokkibegushidoktor/test1/internal/tech/closer"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoClient(ctx context.Context, cfg *config.Config) *mongo.Client {
	opts := options.Client().
		ApplyURI(cfg.MngUri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal().Msgf("Error connecting to mongodb, err: %v", err)
	}

	go func() {
		t := time.NewTicker(cfg.MngPingInterval)
		for range t.C {
			if err = client.Ping(ctx, nil); err != nil {
				log.Error().Msgf("Error pinging mongo, err: %v", err)
			}
		}
	}()

	closer.Add(func() error {
		return client.Disconnect(context.TODO())
	})

	return client
}
