package cmd

import (
	"cobra-sample/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete By Id",
	Long: `The command deletes the name of the item by specified id value`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			fmt.Println("error: invalid id")
			return
		}

		del(id)
	},
}

func init() {
	delCmd.Flags().IntP("id", "i", 0, "Id Field")

	itemsCmd.AddCommand(delCmd)
}

func del(id int) {
	err := storage.Storage.Del(id)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("deleted %v\n", id)
	}
}