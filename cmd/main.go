package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JavaHutt/coinconv/internal/service"
)

var (
	apiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
	isTest = true
)

func main() {
	convSvc, err := service.NewСonverterService(*http.DefaultClient, apiKey, isTest)
	if err != nil {
		log.Fatal(err)
	}
	cmdSvc := service.NewCommandService(convSvc)
	cmdSvc.Exec(os.Args)
}
