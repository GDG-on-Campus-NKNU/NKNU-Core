package schoolbusschedule

import (
	"time"
)

var (
	lastDataFetchTime *time.Time
	ycToHpSchedule    *[]*schedule
	hpToYcSchedule    *[]*schedule
)

func refreshData() error {
	ycToHp, err := fetchYcToHp()
	if err != nil {
		return err
	}
	hpToYc, err := fetchHpToYc()
	if err != nil {
		return err
	}
	ycToHpSchedule = ycToHp
	hpToYcSchedule = hpToYc
	newTime := time.Now()
	lastDataFetchTime = &newTime
	return nil
}
