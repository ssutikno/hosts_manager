package utils

import (
	"log"
	"os"
)

// CheckError is a utility function to handle errors.
func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// GetEnv retrieves the value of the environment variable named by the key.
// It returns the value and a boolean indicating whether the variable was found.
func GetEnv(key string) (string, bool) {
	value, exists := os.LookupEnv(key)
	return value, exists
}