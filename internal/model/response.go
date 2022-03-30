package model

import "time"

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
	Notice       string    `json:"notice"`
}

type Quote struct {
	Price       float32   `json:"price"`
	LsatUpdated time.Time `json:"last_updated"`
}

type DataItem struct {
	ID          int              `json:"id"`
	Symbol      string           `json:"symbol"`
	Name        string           `json:"name"`
	Amount      float32          `json:"amount"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

// SuccessfulResponse is a type for common success response
// from CoinMarketCap API
type SuccessfulResponse struct {
	Status Status     `json:"status"`
	Data   []DataItem `json:"data"`
}

// ErrorResponse is a type for common error response
// from CoinMarketCap API
type ErrorResponse struct {
	Status Status `json:"status"`
}
