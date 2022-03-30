package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JavaHutt/coinconv/internal/service"
	"github.com/JavaHutt/coinconv/utils"
)

func main() {
	apiKey := utils.MustGetEnvString("API_KEY")
	convSvc := service.NewСonverterService(*http.DefaultClient, apiKey)
	value, err := convSvc.Convert("20.1", "USD", "BTC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("converted value: ", value)
}
