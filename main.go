package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const conversionPath = "https://pro-api.coinmarketcap.com/v2/tools/price-conversion"

func mustString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %q is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %q is empty", key)
	}
	return value
}

func main() {
	apiKey := mustString("API_KEY")
	svc := NewService(*http.DefaultClient, apiKey)
	value, err := svc.Convert("20.1", "USD", "BTC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("converted value: ", value)
}

type Service interface {
	Convert(amount, from, to string) (float32, error)
}

type service struct {
	client http.Client
	apiKey string
}

func NewService(client http.Client, apiKey string) *service {
	return &service{client, apiKey}
}

func (s service) Convert(amount, from, to string) (float32, error) {
	if err := validate(amount); err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodGet, conversionPath, nil)
	if err != nil {
		return 0, err
	}
	req = s.withHeaders(req)
	req = s.withQuery(req, amount, from, to)

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, getErrorMessage(resp.Body)
	}
	buff, err := getGzipBody(resp.Body)
	if err != nil {
		return 0, err
	}
	var decodedResponse SuccessfulResponse
	if err = json.Unmarshal(buff, &decodedResponse); err != nil {
		return 0, err
	}
	return decodedResponse.Data[0].Quote[to].Price, nil
}

func (s service) withHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")
	req.Header.Add("X-CMC_PRO_API_KEY", s.apiKey)
	return req
}

func (s service) withQuery(req *http.Request, amount, from, to string) *http.Request {
	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", from)
	q.Add("convert", to)
	req.URL.RawQuery = q.Encode()
	return req
}

func getErrorMessage(body io.ReadCloser) error {
	buff, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	var errorResponse ErrorResponse
	if err = json.Unmarshal(buff, &errorResponse); err != nil {
		return err
	}
	return fmt.Errorf(errorResponse.Status.ErrorMessage)
}

func validate(amount string) error {
	_, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		return fmt.Errorf("bad amount value: %w", err)
	}
	return nil
}

func getGzipBody(body io.ReadCloser) ([]byte, error) {
	reader, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	buff, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

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

type SuccessfulResponse struct {
	Status Status     `json:"status"`
	Data   []DataItem `json:"data"`
}

type ErrorResponse struct {
	Status Status `json:"status"`
}
