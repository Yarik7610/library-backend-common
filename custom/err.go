package custom

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

func NewErr(code int, msg string) *Err {
	return &Err{Code: code, Message: msg}
}
