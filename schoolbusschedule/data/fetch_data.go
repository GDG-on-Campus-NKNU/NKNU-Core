package data

import (
	"encoding/json"
	"io"
	"net/http"
	"nknu-core/schoolbusschedule"
	"strconv"
	"strings"
)

func fetchRawData(url string) (*[]rawScheduleData, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = req.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var data []rawScheduleData
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func fetchData(url string) (*[]*schoolbusschedule.Schedule, error) {
	data, err := fetchRawData(url)
	if err != nil {
		return nil, err
	}

	var schedules []*schoolbusschedule.Schedule
	for _, scheduleData := range *data {
		var s schoolbusschedule.Schedule

		var stations []schoolbusschedule.Station
		for _, stop := range *scheduleData.Stops {
			var sta schoolbusschedule.Station
			sta.Name = stop.Name

			if stop.Note == "教職員工上車" {
				sta.Type = schoolbusschedule.StaffBoarding
			} else if stop.Note == "學生上車" {
				sta.Type = schoolbusschedule.StudentBoarding
			} else if stop.Note == "上車(客滿不停)" {
				sta.Type = schoolbusschedule.BoardingIfNotFull
			} else {
				sta.Type = schoolbusschedule.Alighting
			}

			timeSplit := strings.Split(stop.Time, ":")
			sta.DepartTime.Hour, err = strconv.Atoi(timeSplit[0])
			if err != nil {
				return nil, err
			}
			sta.DepartTime.Minute, err = strconv.Atoi(timeSplit[1])
			if err != nil {
				return nil, err
			}
			stations = append(stations, sta)
		}
		s.Stations = &stations

		// process operates weekdays
		if strings.Contains(scheduleData.Note, "週一～週四加開") {
			s.DaysOfWeek = schoolbusschedule.MondayToThursdayFlag
		} else if strings.Contains(scheduleData.Note, "每天開車") {
			s.DaysOfWeek = schoolbusschedule.AllDaysFlag
		} else if strings.Contains(scheduleData.Note, "週五行駛") {
			s.DaysOfWeek = schoolbusschedule.FridayFlag
		} else {
			s.DaysOfWeek = schoolbusschedule.WeekdayFlag
		}

		if strings.Contains(scheduleData.Note, "例假日停駛") {
			s.OnHoliday = false
		}
		if strings.Contains(scheduleData.Note, "學生專車") {
			s.IsStudentOnly = true
		}

		s.VehicleType = scheduleData.Type
		schedules = append(schedules, &s)
	}
	return &schedules, nil
}

func fetchHpToYc() (*[]*schoolbusschedule.Schedule, error) {
	return fetchData("https://apps.nknu.edu.tw/bus_nosql/toYCJSON")
}

func fetchYcToHp() (*[]*schoolbusschedule.Schedule, error) {
	return fetchData("https://apps.nknu.edu.tw/bus_nosql/toHPJSON")
}
