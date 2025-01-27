package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vlkvch/moneta/internal/fetcher"
)

func mainCurrencies() (string, error) {
	currencies := [...]string{"CNY", "EUR", "KZT", "RUB", "USD"}

	sb := new(strings.Builder)

	for _, curr := range currencies {
		c, err := fetcher.GetCurrency(curr)
		if err != nil {
			return "", err
		}

		fmt.Fprintf(sb, "%-4d %s = %s\n", c.Scale, c.Code, c)
	}

	return strings.Trim(sb.String(), "\n"), nil
}

func singleCurrency(code string, amount float64, quiet bool) (string, error) {
	curr, err := fetcher.GetCurrency(code)
	if err != nil {
		return "", err
	}

	if amount == 0 {
		amount = float64(curr.Scale)
	}

	startStr := fmt.Sprintf("%s %s = ", strconv.FormatFloat(amount, 'f', -1, 32), curr.Code)

	if quiet {
		startStr = ""
	}

	return fmt.Sprintf("%s%s BYN", startStr, strconv.FormatFloat(curr.Rate*amount/float64(curr.Scale), 'f', -1, 32)), nil
}
