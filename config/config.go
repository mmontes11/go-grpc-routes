package config

import (
	"os"
	"strconv"

	"google.golang.org/grpc/testdata"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolVal
}

var (
	// Port indicates the server port
	Port = getEnv("ROUTE_SERVICE_PORT", "11000")
	// TLS indicates if TLS is enabled
	TLS = getBoolEnv("TLS", false)
	// TLScert is the certificate file for TLS
	TLScert = getEnv("TLS_CERT", testdata.Path("server1.pem"))
	// TLSkey is the key file for TLS
	TLSkey = getEnv("TLS_KEY", testdata.Path("server1.key"))
)
