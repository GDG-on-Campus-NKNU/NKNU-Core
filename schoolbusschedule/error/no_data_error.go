package error

type NoScheduleDataError struct {
	Title   string
	Message string
}

func (e *NoScheduleDataError) Error() string {
	return e.Title
}
