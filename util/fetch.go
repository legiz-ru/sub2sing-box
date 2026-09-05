package util

import (
	"io"
	"net/http"
)

func Fetch(url string, maxRetryTimes int, userAgent string) (string, error) {
	retryTime := 0
	var err error
	var resp *http.Response
	for retryTime < maxRetryTimes {
		req, reqErr := http.NewRequest("GET", url, nil)
		if reqErr != nil {
			return "", reqErr
		}
		if userAgent != "" {
			req.Header.Set("User-Agent", userAgent)
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			retryTime++
			continue
		}
		var data []byte
		data, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			retryTime++
			continue
		}
		return string(data), nil
	}
	return "", err
}
