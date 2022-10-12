package error

type FailedQueryValue struct {
	Message string
	Err     error
}

func (uc FailedQueryValue) Error() string {
	return uc.Message
}
