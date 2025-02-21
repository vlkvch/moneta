package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vlkvch/moneta/internal/fetchers"
	"github.com/vlkvch/moneta/internal/models"
)

var (
	amount       = flag.Float64("amount", 0.0, "Set amount to convert.")
	currencyCode = flag.String("from", "RUB", "Set currency to convert from.")
	quiet        = flag.Bool("quiet", false, "Display less output.")
)

const apiURL = "https://api.nbrb.by/exrates/rates"

func init() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage: moneta [option...]\n\nOptions:")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	cacheDir, err := cacheDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "moneta: error: %v\n", models.ErrCacheDirNotFound)
		os.Exit(1)
	}
	os.MkdirAll(cacheDir, 0700)

	fetcher := &fetchers.Fetcher{
		NBRB:  &fetchers.NBRB{ApiURL: apiURL},
		Cache: &fetchers.Cache{CacheDir: cacheDir},
	}

	app := &application{
		fetcher: fetcher,
	}

	var output string

	if len(os.Args[1:]) == 0 {
		main, err := app.mainCurrencies()
		if err != nil {
			fmt.Fprintf(os.Stderr, "moneta: error: %v\n", err)
			os.Exit(1)
		}

		output = main
	} else {
		single, err := app.singleCurrency(*currencyCode, *amount, *quiet)
		if err != nil {
			fmt.Fprintf(os.Stderr, "moneta: error: %v\n", err)
			os.Exit(1)
		}

		output = single
	}

	fmt.Println(output)
}

func cacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(cacheDir, "moneta"), nil
}
