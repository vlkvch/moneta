package fetcher

import (
	"github.com/vlkvch/moneta/internal/cache"
	"github.com/vlkvch/moneta/internal/models"
	"github.com/vlkvch/moneta/internal/nbrb"
)

func GetCurrency(code string) (*models.Currency, error) {
	if cache.IsValid(code) {
		c, err := cache.GetCurrency(code)
		if err != nil {
			return nil, err
		}

		return c, nil
	}

	c, err := nbrb.GetCurrency(code)
	if err != nil {
		return nil, err
	}

	cache.Write(c)

	return c, nil
}
