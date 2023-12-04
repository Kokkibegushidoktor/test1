package config

import "time"

type Config struct {
	MngUri          string        `env:"MNG_URI,required"`
	MngDbName       string        `env:"MNG_DB_NAME,required"`
	MngPingInterval time.Duration `env:"MNG_PING_INTERVAL" envDefault:"10s"`
}
