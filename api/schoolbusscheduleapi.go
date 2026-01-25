package main

import "C"
import (
	"encoding/json"
	"nknu-core/schoolbusschedule"
	"nknu-core/schoolbusschedule/data"
	"nknu-core/utils"
)

func LoadSavedDataApi(toYcData, toHpData string) string {
	if toYcData == "" || toHpData == "" {
		return utils.FormatBase64Output("", schoolbusschedule.NoDataError)
	}
	var toYcSchedule []*schoolbusschedule.Schedule
	err := json.Unmarshal([]byte(toYcData), &toYcSchedule)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	var toHpSchedule []*schoolbusschedule.Schedule
	err = json.Unmarshal([]byte(toHpData), &toHpSchedule)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	data.HpToYcSchedule = &toYcSchedule
	data.YcToHpSchedule = &toHpSchedule
	return utils.FormatBase64Output("", nil)
}

func RefreshSchoolBusDataApi() string {
	err := data.RefreshData()
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output("", nil)
}

func GetLastSchoolBusDataFetchTimeApi() string {
	time, err := data.GetLastSchoolBusDataFetchTime()
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	dataBytes, err := json.Marshal(*time)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetYcToHpScheduleApi() string {
	if data.YcToHpSchedule == nil {
		return utils.FormatBase64Output("", schoolbusschedule.NoDataError)
	}
	dataBytes, err := json.Marshal(*data.YcToHpSchedule)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetHpToYcScheduleApi() string {
	if data.HpToYcSchedule == nil {
		return utils.FormatBase64Output("", schoolbusschedule.NoDataError)
	}
	dataBytes, err := json.Marshal(*data.HpToYcSchedule)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetYcNextBusNowApi() string {
	return getNextBusNow(data.YcToHpSchedule)
}

func GetHpNextBusNowApi() string {
	return getNextBusNow(data.HpToYcSchedule)
}

func getNextBusNow(schedules *[]*schoolbusschedule.Schedule) string {
	index, sche, err := schoolbusschedule.GetNextBusNow(schedules)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	result := make(map[string]interface{})
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetYcNextBusApi(year, month, day, hour, minute int) string {
	return getNextBusApi(data.YcToHpSchedule, year, month, day, hour, minute)
}

func GetHpNextBusApi(year, month, day, hour, minute int) string {
	return getNextBusApi(data.HpToYcSchedule, year, month, day, hour, minute)
}

func getNextBusApi(schedule *[]*schoolbusschedule.Schedule, year, month, day, hour, minute int) string {
	index, sche, err := schoolbusschedule.GetNextBus(schedule, year, month, day, hour, minute)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	result := make(map[string]interface{})
	result["index"] = index
	result["schedule"] = sche
	dataBytes, err := json.Marshal(result)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}

func GetHpBusByIndexApi(index int) string {
	sche, err := schoolbusschedule.GetBusByIndex(data.YcToHpSchedule, index)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	dataBytes, err := json.Marshal(sche)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), err)
}

func GetYcBusByIndexApi(index int) string {
	sche, err := schoolbusschedule.GetBusByIndex(data.YcToHpSchedule, index)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	dataBytes, err := json.Marshal(sche)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), err)
}
