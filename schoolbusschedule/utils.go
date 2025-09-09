package schoolbusschedule

import (
	error2 "NKNU-Core/schoolbusschedule/error"
	"time"
)

func weekdayFlagConvert(weekday time.Weekday) (uint8, error) {
	var weekdays = []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}

	for index, w := range weekdays {
		if w == weekday {
			return 1 << index, nil
		}
	}

	return 0, &error2.WeekdayConvertError{
		Title:   "weekday convert error",
		Message: "weekday convert error",
	}
}
