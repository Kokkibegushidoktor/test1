package models

import "errors"

var (
	ErrTickerAlreadyExists = errors.New("ticker already exists")
	ErrTickerNotFound      = errors.New("ticker not found")
	ErrRatesNotFound       = errors.New("rates not found")
)
