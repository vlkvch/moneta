package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/vlkvch/moneta/internal/fetcher"
)

func mainCurrencies() (string, error) {
	var (
		currencies = [...]string{
			"CNY",
			"EUR",
			"KZT",
			"RUB",
			"USD",
		}
		b strings.Builder
	)

	for i := range currencies {
		c, err := fetcher.GetCurrency(currencies[i])
		if err != nil {
			return "", err
		}

		fmt.Fprintf(&b, "%-4d %s = %s\n", c.Scale, c.Code, c)
	}

	return strings.Trim(b.String(), "\n"), nil
}

func singleCurrency(currencyCode string, amount float64, quiet bool) (string, error) {
	c, err := fetcher.GetCurrency(currencyCode)
	if err != nil {
		return "", err
	}

	if amount == 0 {
		amount = float64(c.Scale)
	}

	startStr := fmt.Sprintf("%s %s = ", strconv.FormatFloat(amount, 'f', -1, 32), c.Code)

	if quiet {
		startStr = ""
	}

	return fmt.Sprintf("%s%s BYN", startStr, strconv.FormatFloat(c.Rate*amount/float64(c.Scale), 'f', -1, 32)), nil
}

func usage() {
	fmt.Println("Usage: moneta [option...]\n\nOptions:")
	flag.PrintDefaults()
}
