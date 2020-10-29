package util

import (
	"log"
	"os"
	"strconv"
)

func GetEnvInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Fatalf("Unable to convert environemt variable %v=%v to in64.")
		}
		return i
	}
	return fallback
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

