package errs

type EndpointError struct {
	Message string `json:"message"`
}

func (e *EndpointError) Error() string {
	return e.Message
}
