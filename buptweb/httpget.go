package buptweb

import (
	"io"
	"net/http"
	"net/url"
)

func HttpGet(turl string, params, headers map[string]string) (string, error) {
	furl := turl
	if params != nil {
		u, err := url.Parse(turl)
		if err != nil {
			return "", err
		}
		uparams := u.Query()
		for k, v := range params {
			uparams.Add(k, v)
		}
		u.RawQuery = uparams.Encode()
		furl = u.String()
	}
	req, err := http.NewRequest(http.MethodGet, furl, nil)
	if err != nil {
		return "", err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
