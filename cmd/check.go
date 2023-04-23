/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bupt-web-cli/buptweb"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check network status",
	Run: func(_ *cobra.Command, _ []string) {
		buptweb.Check()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
