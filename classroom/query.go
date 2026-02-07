package classroom

import (
	"strings"
)

func QueryClassroom(input string) (*Info, error) {
	input = strings.ToUpper(strings.TrimSpace(input))
	if len(input) < 2 {
		return nil, nil
	}

	var prefix string
	var buildingName string
	var floor string
	var classroomID string

	if len(input) >= 3 {
		prefix = input[:2]
		if name, ok := buildingMap[prefix]; ok {
			buildingName = name
			floor = string(input[2])
			classroomID = input[3:]
		}
	}

	if buildingName == "" {
		prefix = input[:1]
		if name, ok := buildingMap[prefix]; ok {
			buildingName = name
			floor = string(input[1])
			classroomID = input[2:]
		}
	}

	if buildingName == "" {
		return nil, nil
	}

	var MapURL string

	hasMap := HasMap(prefix, floor)
	if hasMap {
		MapURL = "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/" + prefix + floor + ".png"
	}

	return &Info{
		BuildingName: buildingName,
		Floor:        floor,
		ClassroomID:  classroomID,
		MapURL:       MapURL,
	}, nil
}
