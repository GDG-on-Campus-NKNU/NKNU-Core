package sso

import (
	"net/url"
	"strings"
)

type Session struct {
	AspNETSessionId string `json:"aspNETSessionId"`
	ViewState       string `json:"viewState"`
}

func Login(session *Session, account string, password string) error {
	loginBody := url.Values{}
	loginBody.Set("uLoginID", account)
	loginBody.Set("uPassword", password)
	loginBody.Set("uLoginPassAuthorizationCode", "登入")
	loginBody.Set("__VIEWSTATE", session.ViewState)
	_, err := newRequest("POST",
		"https://sso.nknu.edu.tw/userLogin/login.aspx?cUrl=%2fdefault.aspx",
		strings.NewReader(loginBody.Encode()),
		session.AspNETSessionId,
		&[]header{
			{"Content-Type", "application/x-www-form-urlencoded"},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
