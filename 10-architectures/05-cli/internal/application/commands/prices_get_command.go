package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"cli-sample/internal/application/services"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Price By Symbol",
	Long:  `The command gets the price of the symbol`,
	Run: func(cmd *cobra.Command, args []string) {
		symbol, err := cmd.Flags().GetString("symbol")
		if err != nil || len(symbol) == 0 {
			fmt.Println("error: invalid symbol")
			return
		}

		getPrice(symbol)
	},
}

func init() {
	getCmd.Flags().StringP("symbol", "s", "", "Symbol")

	pricesCmd.AddCommand(getCmd)
}

func getPrice(symbol string) {
	service := services.NewPriceService()
	price, err := service.GetPrice(symbol)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: %v\n", price.Symbol, price.Price)
}
