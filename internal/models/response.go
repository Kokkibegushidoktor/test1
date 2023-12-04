package models

type FetchResponse struct {
	Ticker     string  `json:"ticker"`
	Price      float32 `json:"price"`
	Difference float32 `json:"difference"`
}
