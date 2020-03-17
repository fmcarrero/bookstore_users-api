package exceptions

type UserNotFound struct {
	ErrMessage string
}

func (e UserNotFound) Error() string {
	return e.ErrMessage
}
