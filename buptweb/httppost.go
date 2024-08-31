package buptweb

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func HttpPostUrlEncoded(turl string, params, headers, data map[string]string, followredirect bool) (string, int, error) {
	furl := turl
	if params != nil {
		u, err := url.Parse(turl)
		if err != nil {
			return "", 0, err
		}
		uparams := u.Query()
		for k, v := range params {
			uparams.Add(k, v)
		}
		u.RawQuery = uparams.Encode()
		furl = u.String()
	}
	var req *http.Request
	var err error
	if data != nil {
		DataUrlVal := url.Values{}
		for k, v := range data {
			DataUrlVal.Add(k, v)
		}
		req, err = http.NewRequest(http.MethodPost, furl, strings.NewReader(DataUrlVal.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest(http.MethodPost, furl, nil)
	}
	if err != nil {
		return "", 0, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if followredirect {
		client.CheckRedirect = nil
	} else {
		client.CheckRedirect = noredirectfunc
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}
	return string(body), resp.StatusCode, nil
}
