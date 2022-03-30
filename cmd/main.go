package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/JavaHutt/coinconv/internal/model"
	"github.com/JavaHutt/coinconv/utils"
)

const conversionPath = "https://pro-api.coinmarketcap.com/v2/tools/price-conversion"

func main() {
	apiKey := utils.MustGetEnvString("API_KEY")
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
	buff, err := utils.ReadGzipBody(resp.Body)
	if err != nil {
		return 0, err
	}
	var decodedResponse model.SuccessfulResponse
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
	var errorResponse model.ErrorResponse
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
