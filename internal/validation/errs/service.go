package errs

type ServiceError struct {
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

type ResourceNotFoundError struct {
	Message string `json:"message"`
}

func (e *ResourceNotFoundError) Error() string {
	return e.Message
}

type NotModifiedError struct {
	Message string `json:"message"`
}

func (e *NotModifiedError) Error() string {
	return e.Message
}
