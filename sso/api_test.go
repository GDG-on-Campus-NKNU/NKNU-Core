package sso

import (
	"os"
	"strings"
	"testing"
)

var (
	account  string
	password string
)

func TestMain(m *testing.M) {
	account = os.Getenv("account")
	password = os.Getenv("password")
	code := m.Run()
	os.Exit(code)
}

func TestWorkflow(t *testing.T) {
	loginSession, err := GetSessionInfo()
	if err != nil {
		t.Fatal(err)
	}

	err = Login(loginSession, account, password)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("GetMailServiceAccount", func(t *testing.T) {
		mailService, err := GetMailServiceAccount(loginSession.AspNETSessionId)
		if err != nil {
			t.Error(err)
		}
		if !strings.HasSuffix(mailService.Google.Account, "@mail.nknu.edu.tw") {
			t.Error("Invalid google account format: " + mailService.Google.Account)
		}
		if !strings.HasSuffix(mailService.O365.Account, "@o365.nknu.edu.tw") {
			t.Error("Invalid o365 account format: " + mailService.O365.Account)
		}
	})

	t.Run("GetHistoryScore", func(t *testing.T) {
		_, err = GetHistoryScore(loginSession)
		if err != nil {
			t.Error(err)
			return
		}
	})
}
