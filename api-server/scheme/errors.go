package scheme

import "fmt"

type ErrorMessage struct {
	Message   string `json:"message"`
	ErrorCode string `json:"code"`
}

func (e *ErrorMessage) InternalServerError(code int, errorMessage string) {
	e.ErrorCode = fmt.Sprintf("ERR_%d", code)
	e.Message = fmt.Sprintf("There was an internal error %s\n", errorMessage)
}

func (e *ErrorMessage) ParsingError(errorMessage string) {
	e.ErrorCode = "ERR_599"
	e.Message = fmt.Sprintf("%s", errorMessage)
}

func (e *ErrorMessage) MethodNotSupport(code int, method string) {
	e.ErrorCode = fmt.Sprintf("ERR_%d", code)
	e.Message = fmt.Sprintf("Http Method %s not allowed.\n", method)
}
