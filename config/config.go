package config

import "os"

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

var (
	// Port indicates the server port
	Port = getEnv("ROUTE_SERVICE_PORT", "11000")
)
