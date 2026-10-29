package config

import "os"

func GetHTTP_ADDR() string {
	address := os.Getenv("HTTP_ADDR")
	if address == "" {
		address = ":8080"
	}
	return address
}
