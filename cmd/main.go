package main

import (
	"log"
	"net/http"
	"os"

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
	cmdSvc := service.NewCommandService(convSvc)
	cmdSvc.Exec(os.Args)
}
