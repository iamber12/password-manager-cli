package serve

import (
	"fmt"
	"github.com/spf13/cobra"
	passwordmanager "v1/internal"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add password",
	Long:  `Add password`,
	Run:   handleAddUpdate,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Add password",
	Long:  `Add password`,
	Run:   handleAddUpdate,
}

func handleAddUpdate(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	site, _ := cmd.Flags().GetString("site")

	if site == "" || username == "" || password == "" {
		fmt.Println("Description, username, and password must be provided.")
		return
	}

	entry := &passwordmanager.PMEntry{"", username, password, site}
	var err error

	if cmd.Name() == "add" {
		err = pm.Add(entry)
	} else {
		err = pm.Update(entry)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Entry saved!")
}

func init() {
	addCmd.Flags().StringP("username", "u", "", "Username for the password")
	addCmd.Flags().StringP("password", "p", "", "Password")
	addCmd.Flags().StringP("site", "s", "", "Site")
	updateCmd.Flags().StringP("username", "u", "", "Username for the password")
	updateCmd.Flags().StringP("password", "p", "", "Password")
	updateCmd.Flags().StringP("site", "s", "", "Site")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
}
