package error

type InternalError struct {
	httpStatus int
	err        error
}

func NewInternalError(httpStatus int, err error) *InternalError {
	return &InternalError{
		httpStatus: httpStatus,
		err:        err,
	}
}

func (e *InternalError) Error() string {
	return e.err.Error()
}

func (e *InternalError) HttpStatus() int {
	return e.httpStatus
}
