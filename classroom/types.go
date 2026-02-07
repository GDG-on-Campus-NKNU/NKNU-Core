package classroom

import "errors"

type Info struct {
	BuildingName string `json:"building_name"`
	Floor        string `json:"floor"`
	ClassroomID  string `json:"classroom_id"`
	MapURL       string `json:"map_url,omitempty"`
}

var UnknnowClassroom = errors.New("unknown classroom")
