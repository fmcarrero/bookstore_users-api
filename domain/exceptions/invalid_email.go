package exceptions

type InvalidEmail struct {
	ErrMessage string
}

func (e InvalidEmail) Error() string {
	return e.ErrMessage
}

func (e InvalidEmail) IsBusinessLogic() bool {
	return true
}
