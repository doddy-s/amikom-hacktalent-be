package errors

func Forbidden(message string) *Error {
	response := new(Error)
	response.Status = "Forbidden"
	response.Message = message
	return response
}