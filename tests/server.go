package tests

import (
	server "gororoba/internal"
	"gororoba/internal/config"
)

func NewTestServer() server.AppServer {
	appConfig := config.Configuration{
		AppConfig: config.AppConfig{
			Name:        "gororoba test",
			Version:     "1.0.0",
			Environment: "test",
		},
		WebConfig: config.WebConfig{
			Port:                     8080,
			IdleTimeout:              10,
			ReadTimeout:              10,
			WriteTimeout:             10,
			ShutdownTimeout:          10,
			GracefulShutdownDisabled: true,
		},
		AWSConfig: config.AWSConfig{
			Region: "us-east-1",
			DynamoDBConfig: config.DynamoDBConfig{
				Endpoint: "http://localhost:4566",
			},
		},
	}

	return server.NewAppServer(appConfig)
}
