package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vlkvch/moneta/internal/cache"
)

var (
	amount       = flag.Float64("amount", 0.0, "Set amount to convert.")
	currencyCode = flag.String("from", "RUB", "Set currency to convert from.")
	quiet        = flag.Bool("quiet", false, "Display less output.")
)

func init() {
	os.MkdirAll(cache.CacheDir(), 0700)
	usage := `Usage: moneta [option...]

Options:
  -amount	Set the amount to convert
  -from		Set the currency to convert from (default RUB)
  -quiet	Display less output`
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), usage)
	}
	flag.Parse()
}

func main() {
	var output string

	if len(os.Args[1:]) == 0 {
		mainCurrencies, err := mainCurrencies()
		if err != nil {
			fmt.Fprintf(os.Stderr, "moneta: error: %v\n", err)
			os.Exit(1)
		}

		output = mainCurrencies
	} else {
		currency, err := singleCurrency(*currencyCode, *amount, *quiet)
		if err != nil {
			fmt.Fprintf(os.Stderr, "moneta: error: %v\n", err)
			os.Exit(1)
		}

		output = currency
	}

	fmt.Println(output)
}
