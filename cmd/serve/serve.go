package serve

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short: "Password Manager CLI Tool",
	Long:  `A simple password manager CLI tool to manage your passwords securely.`,
}

func serve() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initDB)
}

func initDB() {
	// TO DO
}
