/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bupt-web-cli/buptweb"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	login_username string
	login_password string
	login_config   string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login campus network",
	Long: `either call "bupt-web-cli login --username YOUR_USERNAME --password YOUR_PASSWORD"
or "bupt-web-cli login --config SECRETS_JSON_FILENAME"
or "bupt-web-cli login", which will search config.yaml in working dir to login.`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if len(login_username) != 0 && len(login_password) != 0 {
			log.Println("reading username and password from command line arguments")
			buptweb.SetNameAndPass(login_username, login_password)
		} else {
			log.Printf("reading username and password from yaml \"%s\"", login_config)
			buf, err := os.ReadFile(login_config)
			if err != nil {
				log.Fatalln(err)
			}
			err = buptweb.ParseConfig(buf)
			if err != nil {
				log.Fatalln(err)
			}
		}
		buptweb.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&login_username, "username", "u", "", "username to login, typically your student id")
	loginCmd.Flags().StringVarP(&login_password, "password", "p", "", "password to login, typically your password of campus network")
	loginCmd.Flags().StringVarP(&login_config, "config", "c", buptweb.DefaultConfigPath, "should be a yaml file with 'username' and 'password' inside, values are string")
}
