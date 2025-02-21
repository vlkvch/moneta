package models

import "errors"

var (
	ErrNoSuchCurrency   = errors.New("no such currency")
	ErrCacheDirNotFound = errors.New("cache directory not found")
)
