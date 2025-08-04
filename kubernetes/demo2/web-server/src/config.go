package main

import "os"

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func getEnvVar(varName string) string {
	return os.Getenv(varName)
}

func getConfig() Config {
	return Config{
		DBUser:     getEnvVar("DB_USER"),
		DBPassword: getEnvVar("DB_PASSWORD"),
		DBHost:     getEnvVar("DB_HOST"),
		DBPort:     getEnvVar("DB_PORT"),
		DBName:     getEnvVar("DB_NAME"),
	}
}
