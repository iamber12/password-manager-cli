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
	Run:   handleAdd,
}

func handleAdd(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	site, _ := cmd.Flags().GetString("site")

	if site == "" || username == "" || password == "" {
		fmt.Println("Description, username, and password must be provided.")
		return
	}

	user := &passwordmanager.PMUser{"", username, password, site}
	err := pm.Add(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Entry saved successfully!")
}

func init() {
	addCmd.Flags().StringP("username", "u", "", "Username for the password")
	addCmd.Flags().StringP("password", "p", "", "Password")
	addCmd.Flags().StringP("site", "s", "", "Site")
	rootCmd.AddCommand(addCmd)
}
