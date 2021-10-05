package cmd

import (
	"cobra-sample/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set by Id and Name",
	Long: `The command sets the name of the item by specified id value`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			fmt.Println("error: invalid id")
			return
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil || len(name) == 0 {
			fmt.Println("error: invalid name")
			return
		}

		set(id, name)
	},
}

func init() {
	setCmd.Flags().IntP("id", "i", 0, "Id Field")
	setCmd.Flags().StringP("name", "n", "", "Name Field")

	itemsCmd.AddCommand(setCmd)
}

func set(id int, name string) {
	err := storage.Storage.Set(id, name)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("set %v %s\n", id, name)
	}
}