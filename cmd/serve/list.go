package serve

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all passwords",
	Long:  `List all passwords`,
	Run:   handleList,
}

func handleList(cmd *cobra.Command, args []string) {
	entries, err := pm.List()
	if err != nil {
		fmt.Println(err)
		return
	}

	printRecords(entries)
}

func init() {
	rootCmd.AddCommand(listCmd)
}
