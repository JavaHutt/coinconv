package service

import (
	"fmt"
	"strconv"
	"strings"
)

const amountLimit = 1000000000000.0

type converter interface {
	Convert(amount, from, to string) (float32, error)
}

type CommandService struct {
	converterService converter
}

func NewCommandService(converterService converter) *CommandService {
	return &CommandService{converterService}
}

func (s CommandService) Exec(arguments []string) {
	if len(arguments) < 4 {
		writeStandardMessage()
		return
	}
	amount := arguments[1]
	if err := validateAmount(amount); err != nil {
		fmt.Println(err.Error())
		return
	}
	from := arguments[2]
	if err := validateCurrency(from); err != nil {
		fmt.Println(err.Error())
		return
	}
	to := arguments[3]
	s.writeResult(amount, from, to)
}

func (s CommandService) writeResult(amount, from, to string) {
	fmt.Printf("converting %s %s...\n", amount, from)
	list := strings.Split(to, ",")
	if len(list) == 0 {
		fmt.Println("no currencies to convert to")
	}
	for _, v := range list {
		cur := strings.TrimSpace(v)
		if err := validateCurrency(cur); err != nil {
			fmt.Println(err.Error())
			continue
		}
		value, err := s.converterService.Convert(amount, from, cur)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("%s: %f\n", cur, value)
	}
}

func validateAmount(amount string) error {
	parsed, err := strconv.ParseFloat(amount, 10)
	if err != nil {
		return fmt.Errorf("bad amount value: %w", err)
	}
	if parsed > amountLimit {
		return fmt.Errorf("amount must be less than or equal to 1000000000000: %f", amountLimit)
	}
	return nil
}

func validateCurrency(currency string) error {
	if currency == "" {
		return fmt.Errorf("empty currency")
	}
	return nil
}

func writeStandardMessage() {
	fmt.Println("Hi! Use the convertation service like this: ./coinconv 123.45 USD BTC")
	fmt.Println("Where the first argument is the amount, second is a currency that you wish to convert from")
	fmt.Println("And third argument and onward is an array of currencies a previous argument will be converted to.")
	fmt.Println("Use comma to separate these currencies.")
}
