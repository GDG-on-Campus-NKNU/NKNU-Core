package schoolbusschedule

import (
	error2 "NKNU-Core/schoolbusschedule/error"
	"time"
)

func getNextBusNow(schedules *[]*schedule) (int, *schedule, error) {
	tz, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return 0, nil, err
	}
	currentTime := time.Now().In(tz)
	return getNextBus(schedules, currentTime.Year(), int(currentTime.Month()), currentTime.Day(), currentTime.Hour(), currentTime.Minute())
}

func getNextBus(schedules *[]*schedule, year, month, day, hour, minute int) (int, *schedule, error) {
	if schedules == nil {
		return 0, nil, &error2.NoScheduleDataError{
			Title:   "no schedule data",
			Message: "no schedule data",
		}
	}

	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return 0, nil, err
	}
	targetDate := time.Date(year, time.Month(month), day, hour, minute, 0, 0, loc)
	targetWeekdayFlag, err := weekdayFlagConvert(targetDate.Weekday())
	if err != nil {
		return 0, nil, err
	}
	for index, sche := range *schedules {
		for _, stat := range *sche.Stations {
			if ((stat.DepartTime.Hour > targetDate.Hour()) || (stat.DepartTime.Hour == targetDate.Hour() && stat.DepartTime.Minute > targetDate.Minute())) &&
				((targetWeekdayFlag & sche.DaysOfWeek) != 0) &&
				(stat.Type == studentBoarding || stat.Type == staffBoarding || stat.Type == boardingIfNotFull) {
				return index, sche, nil
			}
		}
	}
	return 0, nil, nil
}
