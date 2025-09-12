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
	expectedSchedule     *Schedule
	expectedNextSchedule *Schedule
}

func test(t *testing.T, schedules *[]*Schedule, testCases []testCase) {
	for caseIndex, ca := range testCases {
		index, sche, err := GetNextBus(schedules, ca.year, ca.month, ca.day, ca.hour, ca.minute)
		if err != nil {
			if errors.Is(err, noNextBusError) && ca.expectedSchedule == nil {
				continue
			}
			t.Error(caseIndex, err)
		}
		if sche != ca.expectedSchedule {
			t.Error(caseIndex, "Expected", (*ca.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}

		nextSchedule, err := GetBusByIndex(schedules, index+1)

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
	err := RefreshData()
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = GetNextBusNow(HpToYcSchedule)
	if err != nil && !errors.Is(err, noNextBusError) {
		t.Error(err)
	}
	_, _, err = GetNextBusNow(YcToHpSchedule)
	if err != nil && !errors.Is(err, noNextBusError) {
		t.Error(err)
	}

	t.Run("hpToYcSchedule", func(t *testing.T) {
		test(t, HpToYcSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  3,
				hour:                 7,
				minute:               0,
				expectedSchedule:     (*HpToYcSchedule)[0],
				expectedNextSchedule: (*HpToYcSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  1,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*HpToYcSchedule)[1],
				expectedNextSchedule: (*HpToYcSchedule)[2],
			},
			{
				year:                 2025,
				month:                9,
				day:                  6,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*HpToYcSchedule)[13],
				expectedNextSchedule: nil,
			},
		})
	})

	t.Run("ycToHpSchedule", func(t *testing.T) {
		test(t, YcToHpSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  12,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*YcToHpSchedule)[0],
				expectedNextSchedule: (*YcToHpSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  11,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*YcToHpSchedule)[1],
				expectedNextSchedule: (*YcToHpSchedule)[2],
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
