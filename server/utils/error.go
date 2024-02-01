package utils

type ServerErrorResponse struct {
	ErrorMessage string    `json:"error_message"`
	ErrorCode    ErrorCode `json:"error_code"`
}

type ErrorCode string

const (
	ServerError          = "ServerError"
	EmailAlreadyVerified = "email_already_verified"
)

func GenerateServerErrorResponse(msg string) ServerErrorResponse {
	if msg == "" {
		msg = "unknown internal error"
	}
	return ServerErrorResponse{
		ErrorMessage: msg,
		// ErrorCode:    code,
	}
}
