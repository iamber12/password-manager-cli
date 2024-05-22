package serve

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add password",
	Long:  `Add password`,
	Run:   handleAdd,
}

func handleAdd(cmd *cobra.Command, args []string) {
	description, _ := cmd.Flags().GetString("description")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	if description == "" || username == "" || password == "" {
		fmt.Println("Description, username, and password must be provided.")
		return
	}

	// TODO - logic to handle add in db
}
