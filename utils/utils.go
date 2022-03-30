package utils

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
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

// MustGetEnvString gets the string environment variable
// by key and transforms an error into an exception
func MustGetEnvString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required ENV %s is not set", key)
	}
	if value == "" {
		log.Fatalf("required ENV %s is empty", key)
	}
	return value
}
