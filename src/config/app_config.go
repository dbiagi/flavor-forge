package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

const (
	AppName = "gororoba"
)

type Configuration struct {
	WebConfig
	AppConfig
	DatabaseConfig
	AWSConfig
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

type WebConfig struct {
	Port string
}

type DatabaseConfig struct {
	ConnectionUrl      string
	MaxIdleConnections int
	MaxOpenConnections int
}

type AWSConfig struct {
	Region string
}

func LoadConfig(env string) Configuration {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		slog.Error("Error loading .env file: %v\n", envErr)
		panic(envErr)
	}

	return Configuration{
		WebConfig: WebConfig{
			Port: os.Getenv("PORT"),
		},
		AppConfig: AppConfig{
			Name:        AppName,
			Version:     "1.0.0",
			Environment: env,
		},
		DatabaseConfig: DatabaseConfig{
			ConnectionUrl:      os.Getenv("POSTGRES_URL"),
			MaxIdleConnections: 20,
			MaxOpenConnections: 200,
		},
		AWSConfig: AWSConfig{
			Region: os.Getenv("AWS_REGION"),
		},
	}
}
