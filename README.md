# BUPT Web CLI

Command line interface for login/logout campus network of BUPT. 

This repo is inspired by [Maxlinn/bupt-web-cli](https://github.com/Maxlinn/bupt-web-cli)

## 0x02 Usage

Example 1: login with username and password
`bupt-web-cli login -u YOUR_USERNAME -p YOUR_PASSWORD`

Example 2: login with json file
`bupt-web-cli login -s YOUR_USERNAME_AND_PASSWORD_JSON_FILENAME`

by default, `-s` uses `./secrets.json`. if you follow this manner, just run
`bupt-web-cli login`

Example 3: check login status
`bupt-web-cli check`

Example 4: logout
`bupt-web-cli logout`
