package utils

import "os"

func EnvToString(value *string, key string, defaultValue string) {
	*value = getEnv(key, defaultValue)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
