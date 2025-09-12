package sso

import (
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetSessionInfo() (*Session, error) {
	res, err := http.Get("https://sso.nknu.edu.tw/userLogin/login.aspx?cUrl=/default.aspx")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodyString))
	if err != nil {
		return nil, err
	}
	viewState, exists := doc.Find("#__VIEWSTATE").Attr("value")
	if viewState == "" || !exists {
		return nil, sessionIDNotFoundError
	}
	var sessionId string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "ASP.NET_SessionId" {
			sessionId = cookie.Value
		}
	}
	if sessionId == "" {
		return nil, sessionIDNotFoundError
	}
	return &Session{
		AspNETSessionId: sessionId,
		ViewState:       viewState,
	}, nil
}
