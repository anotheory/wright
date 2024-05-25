package exception

type ValidationException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewValidationException(errorMessage string) *ValidationException {
	return &ValidationException{"0001", errorMessage}
}
