package utils

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
)

// ReadGzipBody reads gzip encodded body
func ReadGzipBody(body io.ReadCloser) ([]byte, error) {
	reader, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	buff, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

// GetString func returns environment variable value as a string value,
// If variable doesn't exist or is not set, returns fallback value
func GetEnvString(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

// GetBool func returns environment variable value as a boolean value,
// If variable doesn't exist or is not set, returns fallback value
func GetBool(key string, fallback bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	if value == "true" || value == "1" {
		return true
	}
	return false
}
