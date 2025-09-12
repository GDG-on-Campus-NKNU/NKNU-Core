package schoolbusschedule

import (
	"time"
)

var (
	LastDataFetchTime *time.Time
	YcToHpSchedule    *[]*Schedule
	HpToYcSchedule    *[]*Schedule
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
