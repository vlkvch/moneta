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
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return ""
	}

	return filepath.Join(cacheDir, "moneta")
}

func IsValid(code string) bool {
	cacheDirFS := os.DirFS(CacheDir())

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

func GetCurrency(code string) (*models.Currency, error) {
	file := currencyCachePath(code)

	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	curr := new(models.Currency)

	if err := json.Unmarshal(contents, curr); err != nil {
		return nil, err
	}

	return curr, nil
}

func Write(curr *models.Currency) error {
	data, err := json.Marshal(curr)
	if err != nil {
		return err
	}

	filePath := currencyCachePath(curr.Code)

	currencyFile, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer currencyFile.Close()

	if _, err := currencyFile.Write(data); err != nil {
		return err
	}

	return nil
}

func currencyCachePath(code string) string {
	return filepath.Join(CacheDir(), strings.ToLower(code)+".json")
}
