package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"cli-sample/internal/infrastructure/config/models"
	"cli-sample/internal/infrastructure/logger/interfaces"
)

var Config models.AppConfig
var Logger interfaces.Logger

var rootCmd = &cobra.Command{
	Use:   "cli-sample",
	Short: "CLI Sample",
	Long:  `The cli sample application that shows prices for crypto`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CLI sample v%s\n", Config.Version)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
