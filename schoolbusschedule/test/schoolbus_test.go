package test

import (
	"errors"
	"nknu-core/schoolbusschedule"
	"nknu-core/schoolbusschedule/data"
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
	expectedSchedule     *schoolbusschedule.Schedule
	expectedNextSchedule *schoolbusschedule.Schedule
}

func test(t *testing.T, schedules *[]*schoolbusschedule.Schedule, testCases []testCase) {
	for caseIndex, ca := range testCases {
		index, sche, err := schoolbusschedule.GetNextBus(schedules, ca.year, ca.month, ca.day, ca.hour, ca.minute)
		if err != nil {
			if errors.Is(err, schoolbusschedule.NoNextBusError) && ca.expectedSchedule == nil {
				continue
			}
			t.Error(caseIndex, err)
		}
		if sche != ca.expectedSchedule {
			t.Error(caseIndex, "Expected", (*ca.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}

		nextSchedule, err := schoolbusschedule.GetBusByIndex(schedules, index+1)

		if err != nil {
			if ca.expectedNextSchedule == nil && errors.Is(err, schoolbusschedule.IndexOutOfRange) {
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
	err := data.RefreshData()
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = schoolbusschedule.GetNextBusNow(data.HpToYcSchedule)
	if err != nil && !errors.Is(err, schoolbusschedule.NoNextBusError) {
		t.Error(err)
	}
	_, _, err = schoolbusschedule.GetNextBusNow(data.YcToHpSchedule)
	if err != nil && !errors.Is(err, schoolbusschedule.NoNextBusError) {
		t.Error(err)
	}

	t.Run("hpToYcSchedule", func(t *testing.T) {
		test(t, data.HpToYcSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  3,
				hour:                 7,
				minute:               0,
				expectedSchedule:     (*data.HpToYcSchedule)[0],
				expectedNextSchedule: (*data.HpToYcSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  1,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*data.HpToYcSchedule)[1],
				expectedNextSchedule: (*data.HpToYcSchedule)[2],
			},
			{
				year:                 2025,
				month:                9,
				day:                  6,
				hour:                 8,
				minute:               0,
				expectedSchedule:     (*data.HpToYcSchedule)[13],
				expectedNextSchedule: nil,
			},
		})
	})

	t.Run("ycToHpSchedule", func(t *testing.T) {
		test(t, data.YcToHpSchedule, []testCase{
			{
				year:                 2025,
				month:                9,
				day:                  12,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*data.YcToHpSchedule)[0],
				expectedNextSchedule: (*data.YcToHpSchedule)[1],
			},
			{
				year:                 2025,
				month:                9,
				day:                  11,
				hour:                 7,
				minute:               10,
				expectedSchedule:     (*data.YcToHpSchedule)[1],
				expectedNextSchedule: (*data.YcToHpSchedule)[2],
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
