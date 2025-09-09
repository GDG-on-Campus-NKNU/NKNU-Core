package sso

import (
	"bytes"
	"io"
	"net/http"
)

type header struct {
	key string
	val string
}

func newRequest(method, url string, body io.Reader, sessionId string, headers *[]header) (string, error) {
	if body == nil {
		body = bytes.NewReader([]byte{})
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("cookie", "ASP.NET_SessionId="+sessionId)
	if headers != nil {
		for _, h := range *headers {
			req.Header.Set(h.key, h.val)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = req.Body.Close()
		_ = res.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	err = isSessionValidate(bodyString)
	if err != nil {
		return "", err
	}
	return bodyString, nil
}
