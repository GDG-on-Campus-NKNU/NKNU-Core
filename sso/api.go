package sso

import "C"
import (
	"NKNU-Core/utils"
	"encoding/json"
)

//export GetSessionInfo
func GetSessionInfo() *C.char {
	session, err := getSessionInfo()
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}

	res, err := json.Marshal(session)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(res), nil))
}

//export Login
func Login(aspNetSessionId *C.char, viewState *C.char, account *C.char, password *C.char) *C.char {
	loginSession := &session{
		AspNETSessionId: C.GoString(aspNetSessionId),
		ViewState:       C.GoString(viewState),
	}
	parsedAccount := C.GoString(account)
	parsedPassword := C.GoString(password)
	err := login(loginSession, parsedAccount, parsedPassword)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output("", nil))
}

//export GetHistoryScore
func GetHistoryScore(aspNetSessionId *C.char, viewState *C.char) *C.char {
	loginSession := &session{
		AspNETSessionId: C.GoString(aspNetSessionId),
		ViewState:       C.GoString(viewState),
	}

	result, err := getHistoryScore(loginSession)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetMailServiceAccount
func GetMailServiceAccount(aspNetSessionId *C.char) *C.char {
	sessionID := C.GoString(aspNetSessionId)
	accounts, err := getMailServiceAccount(sessionID)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	data, err := json.Marshal(accounts)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(data), nil))
}
