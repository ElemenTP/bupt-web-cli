# BUPT Web CLI

Command line interface for login/logout campus network of BUPT.

This repo is inspired by [Maxlinn/bupt-web-cli](https://github.com/Maxlinn/bupt-web-cli)

## 0x01 Usage

Example 1: login with username and password
`bupt-web-cli login -u YOUR_USERNAME -p YOUR_PASSWORD`

---
Example 2: check network status
`bupt-web-cli check`

---
Example 3: login with yaml config file
`bupt-web-cli login -c CONFIG_FILE_PATH`

example config.yaml in example folder
by default, `-c` uses `$HOME/.config/bupt-web-cli/config.yaml`. If you follow this manner, just run
`bupt-web-cli login`

---
Example 4: logout
`bupt-web-cli logout`

---
Example 5: Run as a deamon
`bupt-web-cli login -c CONFIG_FILE_PATH`

config file manner is like `login`
suggest use with systemd service file in example folder
