package retcode

type Error struct {
	ErrCode int
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func NewError(code int, msg string) *Error {
	return &Error{
		ErrCode: code,
		ErrMsg:  msg,
	}
}
