package test

import (
	"encoding/json"
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

func TestGetBusByIndex(t *testing.T) {
	err := data.RefreshData()
	if err != nil {
		t.Fatal(err)
	}

	// Test nil schedules
	_, err = schoolbusschedule.GetBusByIndex(nil, 0)
	if !errors.Is(err, schoolbusschedule.NoDataError) {
		t.Errorf("expected NoDataError, got %v", err)
	}

	// Test index out of range
	_, err = schoolbusschedule.GetBusByIndex(data.HpToYcSchedule, -1)
	if !errors.Is(err, schoolbusschedule.IndexOutOfRange) {
		t.Errorf("expected IndexOutOfRange, got %v", err)
	}

	_, err = schoolbusschedule.GetBusByIndex(data.HpToYcSchedule, len(*data.HpToYcSchedule))
	if !errors.Is(err, schoolbusschedule.IndexOutOfRange) {
		t.Errorf("expected IndexOutOfRange, got %v", err)
	}

	// Test valid index
	sche, err := schoolbusschedule.GetBusByIndex(data.HpToYcSchedule, 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sche != (*data.HpToYcSchedule)[0] {
		t.Error("expected first schedule")
	}
}

func TestGetNextBusErrors(t *testing.T) {
	err := data.RefreshData()
	if err != nil {
		t.Fatal(err)
	}

	// Test nil schedules
	_, _, err = schoolbusschedule.GetNextBus(nil, 2025, 9, 3, 7, 0)
	if !errors.Is(err, schoolbusschedule.NoDataError) {
		t.Errorf("expected NoDataError, got %v", err)
	}

	// Test NoNextBusError
	// HP to YC at 11 PM should have no bus
	_, _, err = schoolbusschedule.GetNextBus(data.HpToYcSchedule, 2025, 9, 3, 23, 0)
	if !errors.Is(err, schoolbusschedule.NoNextBusError) {
		t.Errorf("expected NoNextBusError, got %v", err)
	}
}

func TestGetNextBusVariations(t *testing.T) {
	// Test with a mock schedule to cover different station types
	stations := []schoolbusschedule.Station{
		{
			DepartTime: struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			}{Hour: 8, Minute: 0},
			Name: "Test Station",
			Type: "alighting", // Should be skipped
		},
		{
			DepartTime: struct {
				Hour   int `json:"hour"`
				Minute int `json:"minute"`
			}{Hour: 9, Minute: 0},
			Name: "Boarding Station",
			Type: "studentBoarding", // Should be picked
		},
	}
	mockSchedule := schoolbusschedule.Schedule{
		Stations:   &stations,
		DaysOfWeek: 127, // All days
	}
	schedules := []*schoolbusschedule.Schedule{&mockSchedule}

	// Query at 7:00, should skip alighting at 8:00 and pick studentBoarding at 9:00
	index, _, err := schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if index != 0 {
		t.Error("expected index 0")
	}

	// Test StaffBoarding
	stations[1].Type = "staffBoarding"
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 0)
	if err != nil {
		t.Error("failed staffBoarding")
	}

	// Test BoardingIfNotFull
	stations[1].Type = "boardingIfNotFull"
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 0)
	if err != nil {
		t.Error("failed boardingIfNotFull")
	}

	// Test exact hour, later minute
	stations[1].DepartTime.Hour = 7
	stations[1].DepartTime.Minute = 30
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 0)
	if err != nil {
		t.Error("failed same hour, later minute")
	}

	// Test later hour, earlier minute
	stations[1].DepartTime.Hour = 8
	stations[1].DepartTime.Minute = 0
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 59)
	if err != nil {
		t.Error("failed later hour")
	}

	// Test same time (should NOT match)
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 8, 0)
	if !errors.Is(err, schoolbusschedule.NoNextBusError) {
		t.Error("should not match same time")
	}

	// Test weekday mismatch
	mockSchedule.DaysOfWeek = 64 // Sunday Only
	// Sept 1, 2025 is Monday
	_, _, err = schoolbusschedule.GetNextBus(&schedules, 2025, 9, 1, 7, 0)
	if !errors.Is(err, schoolbusschedule.NoNextBusError) {
		t.Error("should skip if weekday doesn't match")
	}
}

func TestDayFlagsDescription(t *testing.T) {
	testCases := []struct {
		flags    uint8
		expected string
	}{
		{0, "無發車日"},
		{127, "每天"},
		{31, "星期一至星期五"},
		{15, "星期一至星期四"},
		{47, "星期一至星期四、星期六"},
		{95, "星期一至星期五、星期日"},
		{63, "除星期日外每天"},
		{1, "星期一"},
		{64, "星期日"},
		{3, "星期一、星期二"},
	}

	for _, tc := range testCases {
		s := schoolbusschedule.Schedule{
			DaysOfWeek: tc.flags,
			Stations:   &[]schoolbusschedule.Station{},
		}
		data, err := json.Marshal(s)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		var m map[string]interface{}
		json.Unmarshal(data, &m)
		if m["daysOfWeek"] != tc.expected {
			t.Errorf("flags %d: expected %s, got %s", tc.flags, tc.expected, m["daysOfWeek"])
		}
	}
}
