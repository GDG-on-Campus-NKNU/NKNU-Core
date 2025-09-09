package error

type NoLoginError struct {
	Title   string
	Message string
}

func (e *NoLoginError) Error() string {
	return e.Title
}
