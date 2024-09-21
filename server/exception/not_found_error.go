package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) any {
	return NotFoundError{Error: error}
}
