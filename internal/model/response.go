package model

import "time"

// SuccessfulResponse is a type for common success response
// from CoinMarketCap API
type SuccessfulResponse struct {
	Status Status     `json:"status"`
	Data   []DataItem `json:"data"`
}

// GetPrice gets the price value from successful response
func (s SuccessfulResponse) GetPrice(from, to string) float32 {
	return s.Data[0].Quote[to].Price
}

// ErrorResponse is a type for common error response
// from CoinMarketCap API
type ErrorResponse struct {
	Status Status `json:"status"`
}

// SuccessfulResponse is a type for common success response
// from CoinMarketCap Sandbox API
type SuccessfulResponseTest struct {
	Status Status                  `json:"status"`
	Data   map[string]DataItemTest `json:"data"`
}

// GetPrice gets the price value from successful test response
// ... yes, they are different!
func (s SuccessfulResponseTest) GetPrice(from, to string) float32 {
	return s.Data[from].Quote[to].Price
}

// Status is a common type for all responses
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
	LastUpdated time.Time `json:"last_updated"`
}

// Data is a common type for data item
type Data struct {
	Symbol      string           `json:"symbol"`
	Name        string           `json:"name"`
	Amount      float32          `json:"amount"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

// DataItem is a type for a single data item
type DataItem struct {
	ID int `json:"id"`
	Data
}

// DataItemTest is a type for a single data item
// the difference is in ID type
type DataItemTest struct {
	ID string `json:"id"`
	Data
}
