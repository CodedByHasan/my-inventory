package utils

import (
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
func findDotEnv() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get current working directory:", err)
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			log.Printf("Found .env file at %s", envPath)
			return envPath
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached root
		}
		dir = parent
	}

	log.Fatalln(".env file not found in any parent directory")
	return ""
}

func LoadEnvVar(varName string) string {
	envFilePath := findDotEnv()

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
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
