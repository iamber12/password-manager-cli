package serve

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	passwordmanager "v1/internal"
)

var rootCmd = &cobra.Command{
	Short: "Password Manager CLI Tool",
	Long:  `A simple password manager CLI tool to manage your passwords securely.`,
}

var pm = passwordmanager.NewPasswordManagerService()

func Serve() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Invalid command. Use add, update, find, delete or list. Available args: -u for username, -p for password, -s for site")
		os.Exit(1)
	}
}

func printRecords(entries []*passwordmanager.PMEntry) {
	fmt.Printf("%-30s %-20s %-20s\n", "Username", "Password", "Site")
	fmt.Println(strings.Repeat("-", 70))
	for _, entry := range entries {
		fmt.Printf("%-30s %-20s %-20s\n", entry.Username, entry.Password, entry.Site)
	}
}

func init() {
	cobra.OnInitialize()
}
