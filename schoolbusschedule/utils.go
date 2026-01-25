package schoolbusschedule

import (
	"strings"
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

	return 0, flatConvertError
}

func getDayFlagsDescription(flags uint8) string {
	if flags == 0 {
		return "無發車日"
	}
	if flags == AllDaysFlag {
		return "每天"
	}

	var days []string

	if flags&MondayFlag != 0 {
		days = append(days, "星期一")
	}
	if flags&TuesdayFlag != 0 {
		days = append(days, "星期二")
	}
	if flags&WednesdayFlag != 0 {
		days = append(days, "星期三")
	}
	if flags&ThursdayFlag != 0 {
		days = append(days, "星期四")
	}
	if flags&FridayFlag != 0 {
		days = append(days, "星期五")
	}
	if flags&SaturdayFlag != 0 {
		days = append(days, "星期六")
	}
	if flags&SundayFlag != 0 {
		days = append(days, "星期日")
	}

	switch {
	case flags == WeekdayFlag:
		return "星期一至星期五"
	case flags == MondayToThursdayFlag:
		return "星期一至星期四"
	case flags == (MondayToThursdayFlag | SaturdayFlag):
		return "星期一至星期四、星期六"
	case flags == (WeekdayFlag | SundayFlag):
		return "星期一至星期五、星期日"
	case flags == (AllDaysFlag ^ SundayFlag):
		return "除星期日外每天"
	}

	// 如果無法合併，就直接列出
	if len(days) == 0 {
		return "無發車日"
	}
	if len(days) == 1 {
		return days[0]
	}

	// 多天就用「、」分隔
	return strings.Join(days, "、")
}
