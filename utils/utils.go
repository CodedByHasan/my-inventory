package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PathToEnvFile() string {
	value, exists := os.LookupEnv("ENVPATH")

	if exists && value != "" {
		log.Printf("Found ENVPATH variable")
		return value
	} else {
		log.Fatalf("ENVPATH not set. Please set absolute path to .env file")
		return ""
	}
}

func LoadEnvVar(varName string) string {
	envFilePath := PathToEnvFile()

	errReadingFromEnv := godotenv.Load(envFilePath)

	if errReadingFromEnv != nil {
		log.Fatalln("Error loading .env file")
	}

	envVar, doesExists := os.LookupEnv(varName)

	if !doesExists {
		log.Fatalf("%s does not exists. Please ensure it is set in .env file", varName)
	}

	return envVar
}

