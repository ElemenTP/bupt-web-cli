/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bupt-web-cli/buptweb"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var daemon_config string

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "run in daemon mode, periodically check network status",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("reading username and password from yaml \"%s\".\n", daemon_config)
		buf, err := os.ReadFile(daemon_config)
		if err != nil {
			log.Fatalln(err)
		}
		err = buptweb.ParseConfig(buf)
		if err != nil {
			log.Fatalln(err)
		}
		buptweb.Daemon()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().StringVarP(&daemon_config, "config", "c", buptweb.DefaultConfigPath, "should be a yaml file with 'username' and 'password' inside, values are string")
}
