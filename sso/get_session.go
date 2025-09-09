package sso

import (
	error2 "NKNU-Core/sso/error"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getSessionInfo() (*session, error) {
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
		return nil, &error2.SessionIdNotFoundError{
			Title:   "View State not found",
			Message: "View State not found",
		}
	}
	var sessionId string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "ASP.NET_SessionId" {
			sessionId = cookie.Value
		}
	}
	if sessionId == "" {
		return nil, &error2.SessionIdNotFoundError{
			Title:   "Session Id not found",
			Message: "Session Id not found",
		}
	}
	return &session{
		AspNETSessionId: sessionId,
		ViewState:       viewState,
	}, nil
}
