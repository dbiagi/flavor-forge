package main

import (
	"gororoba/cmd/commands"
	"gororoba/config"

	"github.com/spf13/cobra"
)

func main() {
	appConfig := config.LoadConfig("development")
	config.ConfigureLogger(appConfig.AppConfig)

	rootCmd := rootCmd()
	rootCmd.AddCommand(commands.NewCreateRecipesCommand(appConfig.AWSConfig))
	rootCmd.Execute()
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gororoba",
		Short: "Helper to gororoba application",
	}
}
