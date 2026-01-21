package main

import (
	"C"
	"encoding/json"
	"nknu-core/sso"
	ssofuncs "nknu-core/sso/funcs"
	"nknu-core/utils"
)

func GetSessionInfoApi() string {
	session, err := sso.GetSessionInfo()
	if err != nil {
		return utils.FormatBase64Output("", err)
	}

	res, err := json.Marshal(session)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(res), nil)
}

func LoginApi(aspNetSessionId string, viewState string, account string, password string) string {
	loginSession := &sso.Session{
		AspNETSessionId: aspNetSessionId,
		ViewState:       viewState,
	}
	parsedAccount := account
	parsedPassword := password
	err := sso.Login(loginSession, parsedAccount, parsedPassword)
	return utils.FormatBase64Output("", err)
}

func GetHistoryScoreApi(aspNetSessionId string, viewState string) string {
	loginSession := &sso.Session{
		AspNETSessionId: aspNetSessionId,
		ViewState:       viewState,
	}

	result, err := ssofuncs.GetHistoryScore(loginSession)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetMailServiceAccountApi(aspNetSessionId string) string {
	accounts, err := ssofuncs.GetMailServiceAccount(aspNetSessionId)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	data, err := json.Marshal(accounts)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(data), nil)
}
