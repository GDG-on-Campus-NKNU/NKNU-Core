package schoolbusschedule

import (
	"errors"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type testCase struct {
	year                 int
	month                int
	day                  int
	hour                 int
	minute               int
	expectedSchedule     *schedule
	expectedNextSchedule *schedule
}

func test(t *testing.T, schedules *[]*schedule, testCases []testCase) {
	for caseIndex, ca := range testCases {
		index, sche, err := getNextBus(schedules, ca.year, ca.month, ca.day, ca.hour, ca.minute)
		if err != nil {
			if errors.Is(err, noNextBusError) && ca.expectedSchedule == nil {
				continue
			}
			t.Error(caseIndex, err)
		}
		if sche != ca.expectedSchedule {
			t.Error(caseIndex, "Expected", (*ca.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}

		nextSchedule, err := getBusByIndex(schedules, index+1)

		if err != nil {
			if ca.expectedNextSchedule == nil && errors.Is(err, indexOutOfRange) {
				continue
			}
			t.Error(caseIndex, err)
		}
		if nextSchedule != ca.expectedNextSchedule {
			t.Error(caseIndex, "Expected", (*ca.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}
	}
}

func TestWorkflow(t *testing.T) {
	err := refreshData()
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = getNextBusNow(hpToYcSchedule)
	if err != nil {
		t.Error(err)
	}
	_, _, err = getNextBusNow(ycToHpSchedule)
	if err != nil {
		t.Error(err)
	}

	t.Run("hpToYcSchedule", func(t *testing.T) {
		test(t, hpToYcSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  3,
				hour:                 7,
				minute:               0,
				expectedSchedule:     (*hpToYcSchedule)[0],
				expectedNextSchedule: (*hpToYcSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  1,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*hpToYcSchedule)[1],
				expectedNextSchedule: (*hpToYcSchedule)[2],
			},
			{
				year:                 2025,
				month:                9,
				day:                  6,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*hpToYcSchedule)[13],
				expectedNextSchedule: nil,
			},
		})
	})

	t.Run("ycToHpSchedule", func(t *testing.T) {
		test(t, ycToHpSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  12,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*ycToHpSchedule)[0],
				expectedNextSchedule: (*ycToHpSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  11,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*ycToHpSchedule)[1],
				expectedNextSchedule: (*ycToHpSchedule)[2],
			},
			{
				year:                 2025,
				month:                9,
				day:                  13,
				hour:                 7,
				minute:               10,
				expectedSchedule:     nil,
				expectedNextSchedule: nil,
			},
		})
	})
}
