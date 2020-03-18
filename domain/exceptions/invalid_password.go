package exceptions

type InvalidPassword struct {
	ErrMessage string
}

func (e InvalidPassword) Error() string {
	return e.ErrMessage
}
func (e InvalidPassword) IsBusinessLogic() bool {
	return true
}
