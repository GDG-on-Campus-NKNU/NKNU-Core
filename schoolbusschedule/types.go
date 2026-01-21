package schoolbusschedule

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
	return nil, nil
}
