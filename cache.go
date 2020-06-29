package gincachemiddle

import "time"

type Cache interface {
	Get(key string) ([]byte, bool)
	Set(key string, value []byte, expiration time.Duration) error
}
