package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JavaHutt/coinconv/internal/service"
	"github.com/JavaHutt/coinconv/utils"
)

var (
	apiKey = utils.GetEnvString("API_KEY", "")
	isTest = utils.GetBool("IS_TEST", false)
)

func main() {
	convSvc, err := service.NewСonverterService(*http.DefaultClient, apiKey, isTest)
	if err != nil {
		log.Fatal(err)
	}
	value, err := convSvc.Convert("20.1", "USD", "BTC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("converted value: ", value)
}
