package schoolbusschedule

import "errors"

var (
	flatConvertError = errors.New("FlatConvertError")
	noDataError      = errors.New("NoDataError")
	indexOutOfRange  = errors.New("IndexOutOfRange")
	noNextBusError   = errors.New("NoNextBusError")
)
