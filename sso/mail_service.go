package sso

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type mailServiceAccount struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type getMailServiceAccountResponse struct {
	Google *mailServiceAccount `json:"google"`
	O365   *mailServiceAccount `json:"o365"`
}

func getMailServiceAccount(sessionID string) (*getMailServiceAccountResponse, error) {
	googleBodyString, err := newRequest("POST", "https://sso.nknu.edu.tw/Services/GmailandO365.aspx", nil, sessionID, nil)
	if err != nil {
		return nil, err
	}
	o365ReqHeaders := []header{{"Content-Type", "application/x-www-form-urlencoded"}}
	o365ReqBody := url.Values{}
	o365ReqBody.Set("__EVENTARGUMENT", "1")
	o365ReqBody.Set("__EVENTTARGET", "ctl00$phMain$Menu1")
	o365BodyString, err := newRequest("POST", "https://sso.nknu.edu.tw/Services/GmailandO365.aspx", strings.NewReader(o365ReqBody.Encode()), sessionID, &o365ReqHeaders)
	if err != nil {
		return nil, err
	}

	googleDoc, err := goquery.NewDocumentFromReader(strings.NewReader(googleBodyString))
	if err != nil {
		return nil, err
	}
	o365Doc, err := goquery.NewDocumentFromReader(strings.NewReader(o365BodyString))
	if err != nil {
		return nil, err
	}

	googleAccount := googleDoc.Find("label[for=rblGoogleAccount_0]").Text()
	googlePassword := strings.TrimSpace(googleDoc.Find("tr[id=ctl00_phMain_trDefaultPwd]").Find("td").Text())

	o365Account := o365Doc.Find("a[id=ctl00_phMain_hlinko365Account]").Text()
	o365Password := strings.TrimSpace(o365Doc.Find("tr[id=ctl00_phMain_trDefaultPwd2]").Find("td").Text())

	return &getMailServiceAccountResponse{
		Google: &mailServiceAccount{
			Account:  googleAccount,
			Password: googlePassword,
		},
		O365: &mailServiceAccount{
			Account:  o365Account,
			Password: o365Password,
		},
	}, nil
}
