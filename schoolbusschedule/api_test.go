package schoolbusschedule

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type testCase struct {
	year             int
	month            int
	day              int
	hour             int
	minute           int
	expectedSchedule *schedule
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

	for _, toYcTestCase := range []testCase{
		{
			year:             2025,
			month:            9,
			day:              3,
			hour:             7,
			minute:           0,
			expectedSchedule: (*hpToYcSchedule)[0],
		},
		{
			year:             2025,
			month:            9,
			day:              1,
			hour:             8,
			minute:           0,
			expectedSchedule: (*hpToYcSchedule)[1],
		},
		{
			year:             2025,
			month:            9,
			day:              6,
			hour:             8,
			minute:           0,
			expectedSchedule: (*hpToYcSchedule)[13],
		},
	} {
		index, sche, err := getNextBus(hpToYcSchedule, toYcTestCase.year, toYcTestCase.month, toYcTestCase.day, toYcTestCase.hour, toYcTestCase.minute)
		if err != nil {
			t.Error(index, err)
		}
		if sche != toYcTestCase.expectedSchedule {
			t.Error(index, "Expected", (*toYcTestCase.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}
	}

	for _, toHpTestCase := range []testCase{
		{
			year:             2025,
			month:            9,
			day:              12,
			hour:             7,
			minute:           10,
			expectedSchedule: (*ycToHpSchedule)[0],
		},
		{
			year:             2025,
			month:            9,
			day:              11,
			hour:             7,
			minute:           10,
			expectedSchedule: (*ycToHpSchedule)[1],
		},
		{
			year:             2025,
			month:            9,
			day:              13,
			hour:             7,
			minute:           10,
			expectedSchedule: nil,
		},
	} {
		index, sche, err := getNextBus(ycToHpSchedule, toHpTestCase.year, toHpTestCase.month, toHpTestCase.day, toHpTestCase.hour, toHpTestCase.minute)
		if err != nil {
			t.Error(index, err)
		}
		if sche == nil && toHpTestCase.expectedSchedule != nil {
			t.Error(index, "no schedule found")
		}
		if toHpTestCase.expectedSchedule == nil && sche != nil {
			t.Error(index, "expected no schedule found")
		}
		if sche != toHpTestCase.expectedSchedule {
			t.Error(index, "Expected", (*toHpTestCase.expectedSchedule.Stations)[0].DepartTime, "got", (*sche.Stations)[0].DepartTime)
		}
	}
}
