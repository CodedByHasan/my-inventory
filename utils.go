package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnvVar(varName string) string {
	err_reading_from_env := godotenv.Load()

	if err_reading_from_env != nil {
		log.Fatalln("Error loading .env file")
	}

	envVar, doesExists := os.LookupEnv(varName)

	if !doesExists {
		log.Fatalf("%s does not exists. Please ensure it is set in .env file", varName)
	}

	return envVar
}
