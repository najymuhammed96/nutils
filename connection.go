package nutils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func CallURLGET(aurl string, headerMap map[string]string, timeout time.Duration) ([]byte, error) {
	return callURL(aurl, "GET", headerMap, nil, timeout)
}

func CallURLPOST(aurl string, headerMap map[string]string, body []byte, timeout time.Duration) ([]byte, error) {
	return callURL(aurl, "POST", headerMap, body, timeout)
}

func callURL(aurl, method string, headerMap map[string]string, body []byte, timeout time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeout * time.Second}
	req, err := http.NewRequest(method, aurl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Close = true
	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
