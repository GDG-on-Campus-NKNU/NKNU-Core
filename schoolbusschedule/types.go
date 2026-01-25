package schoolbusschedule

import "encoding/json"

type departTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type Station struct {
	DepartTime departTime  `json:"departTime"`
	Name       string      `json:"name"`
	Type       stationType `json:"type"`
}

type Schedule struct {
	Stations      *[]Station `json:"stations"`
	IsStudentOnly bool       `json:"isStudentOnly"`
	OnHoliday     bool       `json:"onHoliday"`
	VehicleType   string     `json:"vehicleType"`
	DaysOfWeek    uint8      `json:"daysOfWeek"` // Bit flags for days this schedule operates
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	parsedResult := make(map[string]interface{})
	parsedResult["stations"] = s.Stations
	parsedResult["isStudentOnly"] = s.IsStudentOnly
	parsedResult["onHoliday"] = s.OnHoliday
	parsedResult["vehicleType"] = s.VehicleType
	parsedResult["daysOfWeek"] = getDayFlagsDescription(s.DaysOfWeek)

	jsonResult, err := json.Marshal(parsedResult)
	if err != nil {
		return nil, err
	}
	return jsonResult, nil
}
