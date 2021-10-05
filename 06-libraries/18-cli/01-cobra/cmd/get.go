package cmd

import (
	"cobra-sample/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get By Id",
	Long: `The command gets the name of the item by specified id value`,
	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Println("error: invalid id")
			return
		}

		if all {
			getAll()
			return
		}

		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			fmt.Println("error: invalid id")
			return
		}

		get(id)
	},
}

func init() {
	getCmd.Flags().BoolP("all", "a", true, "All")
	getCmd.Flags().IntP("id", "i", 0, "Id Field")

	itemsCmd.AddCommand(getCmd)
}

func getAll() {
	items := storage.Storage.GetAll()

	for id, name := range items {
		fmt.Printf("%v - %s\n", id, name)
	}
}

func get(id int) {
	name, err := storage.Storage.Get(id)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%v - %s\n", id, name)
	}
}