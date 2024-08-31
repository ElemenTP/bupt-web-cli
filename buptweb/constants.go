package buptweb

import "net/http"

const PORTAL_BASE_URL string = "http://10.3.8.216"
const WIRED_BASE_URL string = "http://10.3.8.211"
const CONN_TEST_URL string = "http://www.msftconnecttest.com/connecttest.txt"
const CONN_SUCCESS_CODE int = 200
const NAME = "bupt-web-cli"

type Status int

const (
	s_unknown Status = iota
	s_success
	s_failed
	s_noneed
	s_error
)

var noredirectfunc = func(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

var (
	HEADERS map[string]string
	client  *http.Client
)

func init() {
	HEADERS = map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	}
	client = http.DefaultClient
}
