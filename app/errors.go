package app

type AppError struct {
	Status int
	Error error
}

func NewError(status int, err error) *AppError {
	return &AppError{ Status: status, Error: err }
}
