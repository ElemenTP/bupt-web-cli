/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "unknown version"
	BuildTime = "unknown time"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of bupt-web-cli",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("bupt-web-cli %s %s %s with %s %s\n", Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
