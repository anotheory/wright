package exception

type DbErrorException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewDbErrorException(errorMessage string) *DbErrorException {
	return &DbErrorException{"0002", errorMessage}
}
