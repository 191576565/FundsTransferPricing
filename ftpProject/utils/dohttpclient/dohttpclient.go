package dohttpclient

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func DoHttpGet(url string, params map[string]string) (body []byte, err error) {
	hurl := addParams(url, params)
	client := &http.Client{
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get(hurl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func paramsToString(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	return values.Encode()
}

// Add params to a url string.
func addParams(url_ string, params map[string]string) string {
	if len(params) == 0 {
		return url_
	}

	if !strings.Contains(url_, "?") {
		url_ += "?"
	}

	if strings.HasSuffix(url_, "?") || strings.HasSuffix(url_, "&") {
		url_ += paramsToString(params)
	} else {
		url_ += "&" + paramsToString(params)
	}

	return url_
}
