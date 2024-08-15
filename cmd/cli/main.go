package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vlkvch/moneta/internal/cache"
)

var (
	amount       = flag.Float64("amount", 0.0, "Set amount to convert.")
	currencyCode = flag.String("from", "USD", "Set currency to convert from.")
	quiet        = flag.Bool("quiet", false, "Display less output.")
)

func init() {
	os.MkdirAll(cache.CacheDir(), 0700)
	flag.Usage = usage
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
