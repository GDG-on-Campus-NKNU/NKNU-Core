package schoolbusschedule

import (
	"C"
	"NKNU-Core/utils"
	"encoding/json"
)

//export LoadSavedData
func LoadSavedData(toYcData, toHpData *C.char) *C.char {
	if toYcData == nil || toHpData == nil {
		return C.CString(utils.FormatBase64Output("", noDataError))
	}
	var toYcSchedule []*schedule
	err := json.Unmarshal([]byte(C.GoString(toYcData)), &toYcSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	var toHpSchedule []*schedule
	err = json.Unmarshal([]byte(C.GoString(toHpData)), &toHpSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	hpToYcSchedule = &toYcSchedule
	ycToHpSchedule = &toHpSchedule
	return nil
}

//export RefreshSchoolBusData
func RefreshSchoolBusData() *C.char {
	err := refreshData()
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output("", nil))
}

//export GetLastSchoolBusDataFetchTime
func GetLastSchoolBusDataFetchTime() *C.char {
	if lastDataFetchTime == nil {
		return C.CString(utils.FormatBase64Output("", noDataError))
	}
	dataBytes, err := json.Marshal(*lastDataFetchTime)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetYcToHpSchedule
func GetYcToHpSchedule() *C.char {
	if ycToHpSchedule == nil {
		return C.CString(utils.FormatBase64Output("", noDataError))
	}
	dataBytes, err := json.Marshal(*ycToHpSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetHpToYcSchedule
func GetHpToYcSchedule() *C.char {
	if hpToYcSchedule == nil {
		return C.CString(utils.FormatBase64Output("", noDataError))
	}
	dataBytes, err := json.Marshal(*hpToYcSchedule)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), nil))
}

//export GetYcNextBusNow
func GetYcNextBusNow() *C.char {
	index, sche, err := getNextBusNow(ycToHpSchedule)
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
	index, sche, err := getNextBusNow(hpToYcSchedule)
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
	index, sche, err := getNextBus(ycToHpSchedule, year, month, day, hour, minute)
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
	index, sche, err := getNextBus(hpToYcSchedule, year, month, day, hour, minute)
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
	index = int(index)
	sche, err := getBusByIndex(ycToHpSchedule, index)
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
	index = int(index)
	sche, err := getBusByIndex(ycToHpSchedule, index)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	dataBytes, err := json.Marshal(sche)
	if err != nil {
		return C.CString(utils.FormatBase64Output("", err))
	}
	return C.CString(utils.FormatBase64Output(string(dataBytes), err))
}
