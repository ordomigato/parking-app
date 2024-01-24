package utils

type ServerErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func GenerateServerErrorResponse(msg string) ServerErrorResponse {
	if msg == "" {
		msg = "unknown internal error"
	}
	return ServerErrorResponse{
		ErrorMessage: msg,
	}
}
