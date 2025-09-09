package error

type SessionIdNotFoundError struct {
	Title   string
	Message string
}

func (err *SessionIdNotFoundError) Error() string {
	return err.Title
}
