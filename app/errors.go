package app

type AppError struct {
	Message string
	Status int
	Error error
}
