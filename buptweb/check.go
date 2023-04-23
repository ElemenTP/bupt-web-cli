package buptweb

import (
	"log"
	"strings"
)

func Check() {
	state, respbody, err := doCheck()
	if err != nil {
		log.Fatalln(err)
	}
	switch state {
	case s_success:
		log.Println("already logged in")
	case s_noneed:
		log.Println("no need to login")
	case s_failed:
		log.Fatalln("not logged in")
	case s_unknown:
		log.Println("unknown status", "response body:\n", respbody)
	}
}

func doCheck() (Status, string, error) {
	respbody, err := HttpGet(BASE_URL+"/index", nil, HEADERS)
	if err != nil {
		return s_error, "", err
	}
	if strings.Contains(respbody, "您已经登录成功") {
		return s_success, respbody, nil
	} else if strings.Contains(respbody, "您所在的网络无需认证") {
		return s_noneed, respbody, nil
	} else if strings.Contains(respbody, "placeholder=\"输入密码\"") {
		return s_failed, respbody, nil
	} else {
		return s_unknown, respbody, nil
	}
}
