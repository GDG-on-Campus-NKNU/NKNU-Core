package schoolbusschedule

const (
	MondayFlag    uint8 = 1 << 0
	TuesdayFlag   uint8 = 1 << 1
	WednesdayFlag uint8 = 1 << 2
	ThursdayFlag  uint8 = 1 << 3
	FridayFlag    uint8 = 1 << 4
	SaturdayFlag  uint8 = 1 << 5
	SundayFlag    uint8 = 1 << 6

	MondayToThursdayFlag = MondayFlag | TuesdayFlag | WednesdayFlag | ThursdayFlag
	WeekdayFlag          = MondayToThursdayFlag | FridayFlag
	AllDaysFlag          = WeekdayFlag | SaturdayFlag | SundayFlag
)

type stationType string

const (
	StaffBoarding     stationType = "staffBoarding"
	StudentBoarding   stationType = "studentBoarding"
	BoardingIfNotFull stationType = "boardingIfNotFull"
	Alighting         stationType = "alighting"
)
