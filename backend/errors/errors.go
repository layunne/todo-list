package errors

type Error struct {
	Code int   `json:"code"`
	Info string `json:"info"`
}

func New(code int, info string) *Error {
	return &Error{
		Code: code,
		Info: info,
	}
}
