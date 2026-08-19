package env

import "os"

// To get the value by the key it is necessary to
// install the godotenv package and load the envs,
// otherwise it will always use fallback value
func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val;
	}
	return fallback;
}