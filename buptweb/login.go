package buptweb

import (
	"log"
	"strings"
)

func Login(username, password string) {
	state, respbody, err := doLogin(username, password)
	if err != nil {
		log.Fatalln(err)
	}
	switch state {
	case s_success:
		log.Println("login success")
	case s_noneed:
		log.Println("no need to login")
	case s_failed:
		log.Fatalln("login failed, invalid username and password")
	case s_unknown:
		log.Println("unknown status", "response body:\n", respbody)
	}
}

func doLogin(username, password string) (Status, string, error) {
	state, respbody, err := doCheck()
	if err != nil {
		return s_error, "", err
	}
	switch state {
	case s_success:
		return s_success, respbody, nil
	case s_noneed:
		return s_noneed, respbody, nil
	}
	data := map[string]string{
		"user": username,
		"pass": password,
	}
	respbody, err = HttpPostUrlEncoded(BASE_URL+"/login", nil, HEADERS, data)
	if err != nil {
		return s_error, "", err
	}
	if strings.Contains(respbody, "您已经登录成功") {
		return s_success, respbody, nil
	} else if strings.Contains(respbody, "用户名或密码错误 (0x01000004)") {
		return s_failed, respbody, nil
	} else if strings.Contains(respbody, "您所在的网络无需认证") {
		return s_noneed, respbody, nil
	} else {
		return s_unknown, respbody, nil
	}
}
