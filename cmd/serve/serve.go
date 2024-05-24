package serve

import (
	"github.com/spf13/cobra"
	passwordmanager "v1/internal"
)

var rootCmd = &cobra.Command{
	Short: "Password Manager CLI Tool",
	Long:  `A simple password manager CLI tool to manage your passwords securely.`,
}

var pm = passwordmanager.NewPasswordManagerService()

func Serve() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize()
}
