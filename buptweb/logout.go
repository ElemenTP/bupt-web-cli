package buptweb

import (
	"log"
	"strings"
)

func Logout() {
	state, err := doLogout()
	if err != nil {
		log.Fatalln(err)
	}
	switch state {
	case s_success:
		log.Println("logout success")
	case s_noneed:
		log.Println("no need to logout")
	case s_failed:
		log.Fatalln("not logged in")
	case s_unknown:
		log.Println("unknown status")
	}
}

func doLogout() (Status, error) {
	state, err := doCheck()
	if err != nil {
		return s_error, err
	}
	if state == s_failed {
		return s_noneed, nil
	}
	switch config.NetworkType {
	case "auto":
		fallthrough
	case "portal":
		respbody, _, err := HttpGet(PORTAL_BASE_URL+"/logout", nil, HEADERS, true)
		if err != nil {
			return s_error, err
		}
		if strings.Contains(respbody, "您所在的网络无需认证") {
			if config.NetworkType == "portal" {
				return s_noneed, nil
			}
		} else {
			state, err := doCheck()
			if err != nil {
				return s_error, err
			}
			if state == s_failed {
				return s_success, nil
			} else if config.NetworkType == "portal" {
				return s_failed, nil
			}
		}
		fallthrough
	case "wired":
		_, _, err := HttpGet(WIRED_BASE_URL+"/logout", nil, HEADERS, true)
		if err != nil {
			return s_error, err
		}
		state, err := doCheck()
		if err != nil {
			return s_error, err
		}
		if state == s_failed {
			return s_success, nil
		} else {
			return s_failed, nil
		}
	}
	return s_unknown, nil
}
