package models

type FetchResponse struct {
	Ticker     string  `json:"ticker"`
	Price      float64 `json:"price"`
	Difference float64 `json:"difference"`
}
