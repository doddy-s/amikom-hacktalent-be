package errors

func BadRequestError(message string) *Error {
	response := new(Error)
	response.Status = "Bad Request"
	response.Message = message
	return response
}