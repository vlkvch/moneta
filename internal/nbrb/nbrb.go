package nbrb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vlkvch/moneta/internal/models"
)

const apiURL = "https://api.nbrb.by/exrates/rates"

func GetCurrency(code string) (*models.Currency, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s?parammode=2", apiURL, code))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, models.ErrNoSuchCurrency
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c := new(models.Currency)

	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
