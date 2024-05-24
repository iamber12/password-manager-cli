package serve

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "find",
	Short: "Find password",
	Long:  `Find password`,
	Run:   handleFetch,
}

func handleFetch(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	site, _ := cmd.Flags().GetString("site")

	if site == "" && username == "" {
		fmt.Println("Site or username must be provided.")
		return
	}

	entries, err := pm.FindEntries(username, site)

	if err != nil {
		fmt.Println(err)
		return
	}

	printRecords(entries)
}

func init() {
	fetchCmd.Flags().StringP("username", "u", "", "Username for the password")
	fetchCmd.Flags().StringP("site", "s", "", "Site")
	rootCmd.AddCommand(fetchCmd)
}
