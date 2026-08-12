package apperrors

type SystemError int

var (
	ErrApp            SystemError = 1
	ErrUserTerminated SystemError = 130
)
