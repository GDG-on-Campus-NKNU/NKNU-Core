package schoolbusschedule

import (
	"time"
)

func GetNextBusNow(schedules *[]*Schedule) (int, *Schedule, error) {
	tz, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return 0, nil, err
	}
	currentTime := time.Now().In(tz)
	return GetNextBus(schedules, currentTime.Year(), int(currentTime.Month()), currentTime.Day(), currentTime.Hour(), currentTime.Minute())
}

func GetNextBus(schedules *[]*Schedule, year, month, day, hour, minute int) (int, *Schedule, error) {
	if schedules == nil {
		return 0, nil, NoDataError
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
	return 0, nil, noNextBusError
}

func GetBusByIndex(schedules *[]*Schedule, index int) (*Schedule, error) {
	if schedules == nil {
		return nil, NoDataError
	}
	if index < 0 || index >= len(*schedules) {
		return nil, indexOutOfRange
	}
	return (*schedules)[index], nil
}
