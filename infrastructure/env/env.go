package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvFile(devEnv string) string {
	return fmt.Sprintf(".%s.env", devEnv)
}

func getDevEnv() string {
	return os.Getenv("NODE_ENV")
}

func Load() error {
	devEnv := getDevEnv()
	envFile := getEnvFile(devEnv)
	err := godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("failed to load environment file %s: %w", envFile, err)
	}

	return nil
}

func Get(key string) string {
	devEnv := getDevEnv()
	envFile := getEnvFile(devEnv)
	env, err := godotenv.Read(envFile)
	if err != nil {
		log.Fatalf("failed to get environment %s: %q", envFile, err)

		return ""
	}

	value, exists := env[key]
	if !exists {
		log.Printf("Warning: key %s not found in environment variables", key)

		return ""
	}

	return value
}
