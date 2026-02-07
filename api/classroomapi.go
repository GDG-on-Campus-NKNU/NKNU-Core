package main

import "C"
import (
	"encoding/json"
	"nknu-core/classroom"
	"nknu-core/utils"
)

func QueryClassroomApi(input string) string {
	info, err := classroom.QueryClassroom(input)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	if info == nil {
		return utils.FormatBase64Output("", classroom.UnknnowClassroom)
	}

	dataBytes, err := json.Marshal(info)
	if err != nil {
		return utils.FormatBase64Output("", err)
	}
	return utils.FormatBase64Output(string(dataBytes), nil)
}
