package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vlkvch/moneta/internal/models"
)

type NBRB struct {
	ApiURL string
}

func (nbrb *NBRB) GetCurrency(code string) (*models.Currency, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s?parammode=2", nbrb.ApiURL, code))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, models.ErrNoSuchCurrency
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	curr := new(models.Currency)

	if err := json.Unmarshal(data, curr); err != nil {
		return nil, err
	}

	return curr, nil
}
