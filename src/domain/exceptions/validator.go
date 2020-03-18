package exceptions

type Validator interface {
	Error() string
	IsBusinessLogic() bool
}
