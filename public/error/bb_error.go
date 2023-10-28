package error

type BitBridgeError struct {
	Code    int
	Message string
}

func NewBitBridgeError(code int, message string) *BitBridgeError {
	return &BitBridgeError{
		Code:    code,
		Message: message,
	}
}
