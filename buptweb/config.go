package buptweb

import (
	"os"
	P "path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UserName    string `yaml:"username"`
	PassWord    string `yaml:"password"`
	Interval    int    `yaml:"interval"`
	NetworkType string `yaml:"network-type"`
}

func ParseConfig(buf []byte) error {
	tempconfig := &Config{
		UserName:    "default username",
		PassWord:    "default password",
		Interval:    300,
		NetworkType: "auto",
	}
	if err := yaml.Unmarshal(buf, tempconfig); err != nil {
		return err
	}

	if tempconfig.NetworkType != "auto" && tempconfig.NetworkType != "wired" && tempconfig.NetworkType != "portal" {
		tempconfig.NetworkType = "auto"
	}

	config = tempconfig
	return nil
}

func SetNameAndPass(username, password string) {
	config.UserName = username
	config.PassWord = password
}

var (
	config            *Config
	DefaultConfigPath = func() string {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir, _ = os.Getwd()
		}
		homeDir = P.Join(homeDir, ".config", NAME)
		if _, err = os.Stat(homeDir); err != nil {
			if configHome, ok := os.LookupEnv("XDG_CONFIG_HOME"); ok {
				homeDir = P.Join(configHome, NAME)
			}
		}
		return P.Join(homeDir, "config.yaml")
	}()
)

func init() {
	config = &Config{
		UserName:    "default username",
		PassWord:    "default password",
		Interval:    300,
		NetworkType: "auto",
	}
}
