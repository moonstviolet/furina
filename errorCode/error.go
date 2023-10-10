package errorCode

type RespError struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *RespError {
	return &RespError{Code: code, Msg: msg}
}

func (re *RespError) WithErrMsg(errMsg string) *RespError {
	re.Msg += ";" + errMsg
	return re
}

func (re *RespError) AutoErrMsg() *RespError {
	if re.Msg == "" {
		re.Msg = errCodeMsg[re.Code]
	}
	return re
}
