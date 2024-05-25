package exception

type NotFoundException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewNotFoundException(errorMessage string) *NotFoundException {
	return &NotFoundException{"0003", errorMessage}
}
