package schoolbusschedule

type departTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type station struct {
	DepartTime departTime  `json:"departTime"`
	Name       string      `json:"name"`
	Type       stationType `json:"type"`
}

type schedule struct {
	Stations      *[]station `json:"stations"`
	IsStudentOnly bool       `json:"isStudentOnly"`
	OnHoliday     bool       `json:"onHoliday"`
	VehicleType   string     `json:"vehicleType"`
	DaysOfWeek    uint8      `json:"daysOfWeek"` // Bit flags for days this schedule operates
}

func (s schedule) MarshalJSON() ([]byte, error) {
	return nil, nil
}

type rawScheduleData struct {
	Direction string              `json:"direction"`
	Stops     *[]rawScheduleStops `json:"stops"`
	Type      string              `json:"type"`
	Note      string              `json:"note"`
}

type rawScheduleStops struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Note string `json:"note"`
}
