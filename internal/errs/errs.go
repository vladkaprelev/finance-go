package errs

type ErrorCode int

type AppError struct {
	Message string
	Code    ErrorCode
	Err     error
}

// Определение кодов ошибок
const (
	ErrCodeValidation ErrorCode = 1001 // Ошибки валидации
	ErrCodeDatabase   ErrorCode = 2001 // Ошибки базы данных
	ErrCodeNotFound   ErrorCode = 3001 // Ошибки отсутствия данных
	// Добавьте дополнительные коды по необходимости
)

func (err *AppError) Error() string {
	return err.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func newAppError(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewValidationError(message string) *AppError {
	return newAppError(ErrCodeValidation, message, nil)
}

func NewDatabaseError(message string, err error) *AppError {
	return newAppError(ErrCodeDatabase, message, err)
}

func NewNotFoundError(message string, err error) *AppError {
	return newAppError(ErrCodeNotFound, message, err)
}
