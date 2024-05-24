package serve

import (
	"fmt"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete password",
	Long:  `Delete password`,
	Run:   handleDelete,
}

func handleDelete(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	site, _ := cmd.Flags().GetString("site")

	if site == "" || username == "" {
		fmt.Println("Username and site must be provided.")
		return
	}
	
	err := pm.Delete(username, site)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Entry deleted!")
}

func init() {
	delCmd.Flags().StringP("username", "u", "", "Username for the password")
	delCmd.Flags().StringP("site", "s", "", "Site")
	rootCmd.AddCommand(delCmd)
}
