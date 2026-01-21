package data

import "C"
import (
	"nknu-core/schoolbusschedule"
	"time"
)

var (
	LastDataFetchTime *time.Time
	YcToHpSchedule    *[]*schoolbusschedule.Schedule
	HpToYcSchedule    *[]*schoolbusschedule.Schedule
)

func RefreshData() error {
	ycToHp, err := fetchYcToHp()
	if err != nil {
		return err
	}
	hpToYc, err := fetchHpToYc()
	if err != nil {
		return err
	}
	YcToHpSchedule = ycToHp
	HpToYcSchedule = hpToYc
	newTime := time.Now()
	LastDataFetchTime = &newTime
	return nil
}

func GetLastSchoolBusDataFetchTime() (*time.Time, error) {
	if LastDataFetchTime == nil {
		return nil, schoolbusschedule.NoDataError
	}
	return LastDataFetchTime, nil
}
