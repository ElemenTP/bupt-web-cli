package buptweb

import (
	"log"
	"strings"
)

func Logout() {
	state, respbody, err := doLogout()
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
		log.Println("unknown status", "response body:\n", respbody)
	}
}

func doLogout() (Status, string, error) {
	state, respbody, err := doCheck()
	if err != nil {
		return s_error, "", err
	}
	switch state {
	case s_failed:
		return s_failed, respbody, nil
	case s_noneed:
		return s_noneed, respbody, nil
	}
	respbody, err = HttpGet(BASE_URL+"/logout", nil, HEADERS)
	if err != nil {
		return s_error, "", err
	}
	if strings.Contains(respbody, "您所在的网络无需认证") {
		return s_noneed, respbody, nil
	} else if strings.Contains(respbody, "placeholder=\"输入密码\"") {
		return s_success, respbody, nil
	} else {
		return s_unknown, respbody, nil
	}
}
