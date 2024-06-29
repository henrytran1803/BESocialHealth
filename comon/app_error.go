package comon

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"error"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullCodeErrorResponse(statusCode int, rootErr error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootError:  rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}
