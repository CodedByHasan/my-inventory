package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnvVar(varName string) string {
	errReadingFromEnv := godotenv.Load()

	if errReadingFromEnv != nil {
		log.Fatalln("Error loading .env file")
	}

	envVar, doesExists := os.LookupEnv(varName)

	if !doesExists {
		log.Fatalf("%s does not exists. Please ensure it is set in .env file", varName)
	}

	return envVar
}
