package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vlkvch/moneta/internal/models"
)

func CacheDir() string {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return ""
	}

	return filepath.Join(userCacheDir, "moneta")
}

func IsValid(currencyCode string) bool {
	cacheDirFS := os.DirFS(CacheDir())

	f, err := cacheDirFS.Open(strings.ToLower(currencyCode) + ".json")
	if err != nil {
		return false
	}

	fileStat, err := f.Stat()
	if err != nil {
		return false
	}

	lastUpdated := fileStat.ModTime()
	resetTime := time.Date(lastUpdated.Year(), lastUpdated.Month(), lastUpdated.Day()+1, 0, 0, 0, 0, time.Local)

	return time.Now().Before(resetTime)
}

func GetCurrency(currencyCode string) (*models.Currency, error) {
	file := filepath.Join(CacheDir(), strings.ToLower(currencyCode)+".json")

	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	currency := new(models.Currency)

	json.Unmarshal(contents, currency)

	return currency, nil
}

func Write(curr *models.Currency) error {
	data, err := json.Marshal(curr)
	if err != nil {
		return err
	}

	file := filepath.Join(CacheDir(), strings.ToLower(curr.Code)+".json")

	f, err := os.OpenFile(file, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
