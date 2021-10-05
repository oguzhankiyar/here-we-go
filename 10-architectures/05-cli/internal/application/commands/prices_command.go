package commands

import (
	"github.com/spf13/cobra"
)

var pricesCmd = &cobra.Command{
	Use:   "prices",
	Short: "Prices",
	Long:  `Prices.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(pricesCmd)
}
