package errors

const (
	Success = iota
	InternalServerError
	NotFound
	BadRequest
	Unauthorized
	Forbidden
	Conflict
	TooManyRequests
	ServiceUnavailable
	NotImplemented
	Timeout
)

const (
	DatabaseConnectionFailed = iota + 100
	DatabaseQueryFailed
	RecordNotFound
	RecordAlreadyExists
)

const (
	InvalidCredentials = iota + 200
	AccessDenied
	SessionExpired
	TokenInvalid
	TokenExpired
)

const (
	UserCreationFailed = iota + 300
	UserUpdateFailed
	UserDeletionFailed
	PasswordHashingFailed
	EmailAlreadyExists
	UsernameAlreadyExists
)

const (
	InvalidInputData = iota + 400
	MissingRequiredField
	InvalidEmailFormat
	InvalidDateFormat
)

const (
	NetworkError = iota + 500
	ServiceCommunicationFailed
)

const (
	BusinessRuleViolation = iota + 600
	OperationNotAllowed
)

type ServiceError struct {
	Code    int
	Message string
	Cause   error
}

func NewServiceError(code int, message string, cause error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func (e *ServiceError) Error() string {
	return e.Message
}
