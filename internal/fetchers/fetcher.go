package fetchers

import "github.com/vlkvch/moneta/internal/models"

type Fetcher struct {
	NBRB  *NBRB
	Cache *Cache
}

func (f *Fetcher) GetCurrency(code string) (*models.Currency, error) {
	if f.Cache.Valid(code) {
		curr, err := f.Cache.GetCurrency(code)
		if err != nil {
			return nil, err
		}

		return curr, nil
	}

	curr, err := f.NBRB.GetCurrency(code)
	if err != nil {
		return nil, err
	}

	err = f.Cache.Write(curr)
	if err != nil {
		return nil, err
	}

	return curr, nil
}
