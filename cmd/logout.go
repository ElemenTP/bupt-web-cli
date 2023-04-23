/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bupt-web-cli/buptweb"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout campus network",
	Run: func(_ *cobra.Command, _ []string) {
		buptweb.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
