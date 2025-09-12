package schoolbusschedule

import "errors"

var (
	flatConvertError = errors.New("FlatConvertError")
	NoDataError      = errors.New("NoDataError")
	indexOutOfRange  = errors.New("IndexOutOfRange")
	noNextBusError   = errors.New("NoNextBusError")
)
