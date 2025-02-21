package fetchers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vlkvch/moneta/internal/models"
)

type Cache struct {
	CacheDir string
}

func (c *Cache) Valid(code string) bool {
	cacheDirFS := os.DirFS(c.CacheDir)

	currencyFile, err := cacheDirFS.Open(strings.ToLower(code) + ".json")
	if err != nil {
		return false
	}
	defer currencyFile.Close()

	fileStat, err := currencyFile.Stat()
	if err != nil {
		return false
	}

	lastUpdated := fileStat.ModTime()
	resetTime := time.Date(lastUpdated.Year(), lastUpdated.Month(), lastUpdated.Day()+1, 0, 0, 0, 0, time.Local)

	return time.Now().Before(resetTime)
}

func (c *Cache) GetCurrency(code string) (*models.Currency, error) {
	cachePath := c.currencyCachePath(code)

	contents, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, err
	}

	curr := new(models.Currency)

	if err := json.Unmarshal(contents, curr); err != nil {
		return nil, err
	}

	return curr, nil
}

func (c *Cache) Write(curr *models.Currency) error {
	data, err := json.Marshal(curr)
	if err != nil {
		return err
	}

	cachePath := c.currencyCachePath(curr.Code)

	currencyFile, err := os.OpenFile(cachePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer currencyFile.Close()

	if _, err := currencyFile.Write(data); err != nil {
		return err
	}

	return nil
}

func (c *Cache) currencyCachePath(code string) string {
	return filepath.Join(c.CacheDir, strings.ToLower(code)+".json")
}
