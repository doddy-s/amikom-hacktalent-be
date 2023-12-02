package errors

func NotFound(message string) *Error {
	response := new(Error)
	response.Status = "Not Found"
	response.Message = message
	return response
}