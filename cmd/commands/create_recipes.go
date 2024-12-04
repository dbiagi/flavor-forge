package commands

import (
	"encoding/json"
	"fmt"
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/repository"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

type cobraCommand func(*cobra.Command, []string)

func NewCreateRecipesCommand(c config.AWSConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "create-recipes",
		Run:       command(c),
		ValidArgs: []string{"file-path"},
	}

	return cmd
}

func command(c config.AWSConfig) cobraCommand {
	return func(cmd *cobra.Command, args []string) {
		dynamoDb, dynamoConnectError := config.CreateDynamoDBConnection(c)
		if dynamoConnectError != nil {
			slog.Error(fmt.Sprintf("Failed to run command %v.", dynamoConnectError))
			return
		}
		r := repository.NewRecipeRepository(dynamoDb)

		jsonFile := args[0]

		slog.Info(fmt.Sprintf("Creating recipes from file %s", jsonFile))

		bytes, readError := os.ReadFile(jsonFile)
		if readError != nil {
			slog.Error(fmt.Sprintf("Failed to read file: %v", readError))
			return
		}

		var recipes []domain.Recipe
		if err := json.Unmarshal(bytes, &recipes); err != nil {
			slog.Error(fmt.Sprintf("Failed to unmarshal JSON: %v", err))
			return
		}

		for _, recipe := range recipes {
			insertRecipe(recipe, r)
		}
	}
}

func insertRecipe(recipe domain.Recipe, r repository.RecipeRepository) {
	if err := r.CreateRecipe(&recipe); err != nil {
		slog.Error(fmt.Sprintf("Failed to create recipe: %s", recipe.Title))
		return
	}
}
