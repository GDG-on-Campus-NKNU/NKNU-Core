package error

type WeekdayConvertError struct {
	Title   string
	Message string
}

func (e *WeekdayConvertError) Error() string {
	return e.Message
}
