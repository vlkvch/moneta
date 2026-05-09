package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (app *application) mainCurrencies() (string, error) {
	currencyCodes := [...]string{"CNY", "EUR", "KZT", "RUB", "USD"}

	sb := new(strings.Builder)

	for _, code := range currencyCodes {
		curr, err := app.fetcher.GetCurrency(code)
		if err != nil {
			return "", err
		}

		fmt.Fprintf(sb, "%4d %s = %s\n", curr.Scale, curr.Code, curr)
	}

	return strings.Trim(sb.String(), "\n"), nil
}

func (app *application) singleCurrency(code string, amount float64, quiet bool) (string, error) {
	curr, err := app.fetcher.GetCurrency(code)
	if err != nil {
		return "", err
	}

	if amount == 0 {
		amount = float64(curr.Scale)
	}

	startString := fmt.Sprintf("%s %s = ", strconv.FormatFloat(amount, 'f', -1, 32), curr.Code)

	if quiet {
		startString = ""
	}

	return fmt.Sprintf("%s%s BYN", startString, strconv.FormatFloat(curr.Rate*amount/float64(curr.Scale), 'f', -1, 32)), nil
}
