package buptweb

import "net/http"

const BASE_URL string = "http://10.3.8.216"

type Status int

const (
	s_error Status = iota
	s_success
	s_failed
	s_noneed
	s_unknown
)

var (
	HEADERS map[string]string
	client  http.Client
)

func init() {
	HEADERS = map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
	}
	client = *http.DefaultClient
}
