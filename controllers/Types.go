package controllers

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newSuccess(data interface{}) *Success {
	response := new(Success)
	response.Status = "Success"
	response.Message = "Success"
	response.Data = data
	return response
}

type Empty struct {
	Nah int8 `json:"status,omitempty"`
}