package buptweb

import (
	"log"
	"strings"
)

func Login() {
	state, err := doLogin()
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
		log.Println("unknown status")
	}
}

func doLogin() (Status, error) {
	state, err := doCheck()
	if err != nil {
		return s_error, err
	}
	switch state {
	case s_success:
		return s_success, nil
	case s_noneed:
		return s_noneed, nil
	}
	data := map[string]string{
		"user": config.UserName,
		"pass": config.PassWord,
	}
	switch config.NetworkType {
	case "auto":
		fallthrough
	case "portal":
		respbody, _, err := HttpPostUrlEncoded(PORTAL_BASE_URL+"/login", nil, HEADERS, data, true)
		if err != nil {
			return s_error, err
		}
		if strings.Contains(respbody, "您已经登录成功") {
			return s_success, nil
		} else if strings.Contains(respbody, "用户名或密码错误") {
			return s_failed, nil
		} else if strings.Contains(respbody, "您所在的网络无需认证") {
			if config.NetworkType == "portal" {
				return s_noneed, nil
			}
		} else {
			return s_unknown, nil
		}
		fallthrough
	case "wired":
		respbody, _, err := HttpPostUrlEncoded(WIRED_BASE_URL+"/login", nil, HEADERS, data, true)
		if err != nil {
			return s_error, err
		}
		if strings.Contains(respbody, "您已经登录成功") {
			return s_success, nil
		} else if strings.Contains(respbody, "用户名或密码错误") {
			return s_failed, nil
		} else if strings.Contains(respbody, "您所在的网络无需认证") {
			return s_noneed, nil
		} else {
			return s_unknown, nil
		}
	}
	return s_unknown, nil
}
