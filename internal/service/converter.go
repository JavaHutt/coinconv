package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"unicode"

	"github.com/JavaHutt/coinconv/internal/model"
	"github.com/JavaHutt/coinconv/utils"
)

const (
	conversionPath = "https://%s-api.coinmarketcap.com/v2/tools/price-conversion"
	testPrefix     = "sandbox"
	publicPrefix   = "pro"
)

type response interface {
	GetPrice(from, to string) float32
}

type ConverterService struct {
	client   http.Client
	apiKey   string
	isTest   bool
	endpoint string
}

func NewСonverterService(client http.Client, apiKey string, isTest bool) (*ConverterService, error) {
	if !isTest && apiKey == "" {
		return nil, fmt.Errorf("you need an api key to use PRO API version")
	}
	endpoint := fmt.Sprintf(conversionPath, publicPrefix)
	if isTest {
		endpoint = fmt.Sprintf(conversionPath, testPrefix)
	}
	return &ConverterService{client, apiKey, isTest, endpoint}, nil
}

func (s ConverterService) Convert(amount, from, to string) (float32, error) {
	req, err := http.NewRequest(http.MethodGet, s.endpoint, nil)
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
	var res response
	if s.isTest {
		if res, err = getResponse[model.SuccessfulResponseTest](buff); err != nil {
			return 0, err
		}
		return res.GetPrice(from, to), nil
	}

	if res, err = getResponse[model.SuccessfulResponse](buff); err != nil {
		return 0, err
	}
	return res.GetPrice(from, to), nil
}

func getResponse[V model.SuccessfulResponse | model.SuccessfulResponseTest](buff []byte) (V, error) {
	var decodedResponse V
	if err := json.Unmarshal(buff, &decodedResponse); err != nil {
		return decodedResponse, err
	}
	return decodedResponse, nil
}

func (s ConverterService) withHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Accept-Encoding", "deflate, gzip")
	req.Header.Add("X-CMC_PRO_API_KEY", s.apiKey)
	return req
}

func (s ConverterService) withQuery(req *http.Request, amount, from, to string) *http.Request {
	convertParam := "convert"
	if unicode.IsDigit(rune(to[0])) {
		convertParam = "convert_id"
	}
	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", from)
	q.Add(convertParam, to)
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
