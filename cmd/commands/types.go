package commands

import (
	"gororoba/controller"

	"github.com/spf13/cobra"
)

type CommandFunction func(*cobra.Command, []string)

type Controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}
