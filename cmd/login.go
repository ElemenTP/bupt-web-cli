/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bupt-web-cli/buptweb"
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	username string
	password string
	secrets  string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login campus network",
	Long: `Either call "bupt-web-cli login --username YOUR_USERNAME --password YOUR_PASSWORD"
or "bupt-web-cli login --secrets SECRETS_JSON_FILENAME"
or "bupt-web-cli login", which will search secrets.json in working dir to login.`,
	Args: cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if len(username) != 0 && len(password) != 0 {
			log.Println("reading username and password from command line arguments")
		} else {
			log.Printf("reading username and password from json \"%s\".\n", secrets)
			jsonFile, err := os.ReadFile(secrets)
			if err != nil {
				log.Fatalln(err)
			}
			j := make(map[string]interface{})
			err = json.Unmarshal(jsonFile, &j)
			if err != nil {
				log.Fatalln(err)
			}
			for k, v := range j {
				if k == "username" {
					switch v := v.(type) {
					case string:
						username = v
					default:
						log.Fatalln("username must be a string")
					}
				} else if k == "password" {
					switch v := v.(type) {
					case string:
						password = v
					default:
						log.Fatalln("password must be a string")
					}
				} else {
					continue
				}
			}
		}
		log.Printf("received username = \"%s\" and password = \"%s\"", username, password)
		if len(username) == 0 || len(password) == 0 {
			log.Fatal("username or password is empty")
		}
		buptweb.Login(username, password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username to login, typically your student id")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password to login, typically your password campus network")
	loginCmd.Flags().StringVarP(&secrets, "secrets", "s", "secrets.json", "Should be a json file with `username` and `password` inside, values are string")
}
