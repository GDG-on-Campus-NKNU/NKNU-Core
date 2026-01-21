package data

type rawScheduleData struct {
	Direction string              `json:"direction"`
	Stops     *[]rawScheduleStops `json:"stops"`
	Type      string              `json:"type"`
	Note      string              `json:"note"`
}
