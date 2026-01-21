package schoolbusschedule

import "errors"

var (
	flatConvertError = errors.New("FlatConvertError")
	NoDataError      = errors.New("NoDataError")
	IndexOutOfRange  = errors.New("IndexOutOfRange")
	NoNextBusError   = errors.New("NoNextBusError")
)
