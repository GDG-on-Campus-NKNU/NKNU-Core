package main

import (
	"C"
	"NKNU-Core/schoolbusschedule"
	"NKNU-Core/sso"
	"NKNU-Core/utils"
	"encoding/json"
)

// SSO

//export GetSessionInfo
func GetSessionInfo() *C.char {
	session, err := sso.GetSessionInfo()
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
	loginSession := &sso.Session{
		AspNETSessionId: C.GoString(aspNetSessionId),
		ViewState:       C.GoString(viewState),
	}
	parsedAccount := C.GoString(account)
	parsedPassword := C.GoString(password)
	err := sso.Login(loginSession, parsedAccount, parsedPassword)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output("", nil))
}

//export GetHistoryScore
func GetHistoryScore(aspNetSessionId *C.char, viewState *C.char) *C.char {
	loginSession := &sso.Session{
		AspNETSessionId: C.GoString(aspNetSessionId),
		ViewState:       C.GoString(viewState),
	}

	result, err := sso.GetHistoryScore(loginSession)
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
	accounts, err := sso.GetMailServiceAccount(sessionID)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	data, err := json.Marshal(accounts)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(data), nil))
}

// school bus schedule

//export LoadSavedData
func LoadSavedData(toYcData, toHpData *C.char) *C.char {
	if toYcData == nil || toHpData == nil {
		return C.CString(utils.FormatBase64Output("", schoolbusschedule.NoDataError))
	}
	var toYcSchedule []*schoolbusschedule.Schedule
	err := json.Unmarshal([]byte(C.GoString(toYcData)), &toYcSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var toHpSchedule []*schoolbusschedule.Schedule
	err = json.Unmarshal([]byte(C.GoString(toHpData)), &toHpSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	schoolbusschedule.HpToYcSchedule = &toYcSchedule
	schoolbusschedule.YcToHpSchedule = &toHpSchedule
	return nil
}

//export RefreshSchoolBusData
func RefreshSchoolBusData() *C.char {
	err := schoolbusschedule.RefreshData()
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output("", nil))
}

//export GetLastSchoolBusDataFetchTime
func GetLastSchoolBusDataFetchTime() *C.char {
	if schoolbusschedule.LastDataFetchTime == nil {
		return C.CString(utils.FormatBase64Output("", schoolbusschedule.NoDataError))
	}
	dataBytes, err := json.Marshal(*schoolbusschedule.LastDataFetchTime)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetYcToHpSchedule
func GetYcToHpSchedule() *C.char {
	if schoolbusschedule.YcToHpSchedule == nil {
		return C.CString(utils.FormatBase64Output("", schoolbusschedule.NoDataError))
	}
	dataBytes, err := json.Marshal(*schoolbusschedule.YcToHpSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetHpToYcSchedule
func GetHpToYcSchedule() *C.char {
	if schoolbusschedule.HpToYcSchedule == nil {
		return C.CString(utils.FormatBase64Output("", schoolbusschedule.NoDataError))
	}
	dataBytes, err := json.Marshal(*schoolbusschedule.HpToYcSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetYcNextBusNow
func GetYcNextBusNow() *C.char {
	index, sche, err := schoolbusschedule.GetNextBusNow(schoolbusschedule.YcToHpSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var result map[string]interface{}
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetHpNextBusNow
func GetHpNextBusNow() *C.char {
	index, sche, err := schoolbusschedule.GetNextBusNow(schoolbusschedule.HpToYcSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var result map[string]interface{}
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetYcNextBus
func GetYcNextBus(rawYear, rawMonth, rawDay, rawHour, rawMinute C.int) *C.char {
	year := int(rawYear)
	month := int(rawMonth)
	day := int(rawDay)
	hour := int(rawHour)
	minute := int(rawMinute)
	index, sche, err := schoolbusschedule.GetNextBus(schoolbusschedule.YcToHpSchedule, year, month, day, hour, minute)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var result map[string]interface{}
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetHpNextBus
func GetHpNextBus(rawYear, rawMonth, rawDay, rawHour, rawMinute C.int) *C.char {
	year := int(rawYear)
	month := int(rawMonth)
	day := int(rawDay)
	hour := int(rawHour)
	minute := int(rawMinute)
	index, sche, err := schoolbusschedule.GetNextBus(schoolbusschedule.HpToYcSchedule, year, month, day, hour, minute)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var result map[string]interface{}
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetHpBusByIndex
func GetHpBusByIndex(index C.int) *C.char {
	parsedIndex := int(index)
	sche, err := schoolbusschedule.GetBusByIndex(schoolbusschedule.YcToHpSchedule, parsedIndex)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	dataBytes, err := json.Marshal(sche)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), err))
}

//export GetYcBusByIndex
func GetYcBusByIndex(index C.int) *C.char {
	parsedIndex := int(index)
	sche, err := schoolbusschedule.GetBusByIndex(schoolbusschedule.YcToHpSchedule, parsedIndex)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	dataBytes, err := json.Marshal(sche)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), err))
}

// main

func main() {

}
