package errs

type DatabaseError struct {
	Message string `json:"error"`
}

func (e *DatabaseError) Error() string {
	return e.Message
}
