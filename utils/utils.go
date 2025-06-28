package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	AppPort    string
}

// walk up directory tree until .env is found
func findDotEnv() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get current working directory:", err)
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			log.Printf("Found .env file at %s", envPath)
			return envPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached root
		}
		dir = parent
	}
	return "", fmt.Errorf(".env file not found in any parent directory")
}

func LoadEnvVar(varName string) string {
	envFilePath, err := findDotEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	val, exists := os.LookupEnv(varName)
	if !exists {
		log.Fatalf("%s not set in .env file", varName)
	}
	return val
}

func LoadEnvConfig() EnvConfig {
	return EnvConfig{
		DBUser:     LoadEnvVar("DB_USER"),
		DBPassword: LoadEnvVar("DB_PASSWORD"),
		DBName:     LoadEnvVar("DB_NAME"),
		DBPort:     LoadEnvVar("DB_PORT"),
		AppPort:    LoadEnvVar("APP_PORT"),
	}
}
