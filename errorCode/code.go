package errorCode

const (
	CODE_OK              = 0
	CODE_SYSTEM_ERROR    = 1
	CODE_PARAM_WRONG     = 2
	CODE_NOT_LOGINED     = 3
	CODE_ILLEGAL_REQUEST = 4
	CODE_NO_PERMISSION   = 5
	CODE_NOT_FOUND       = 6
)

var (
	errCodeMsg = map[int]string{
		CODE_OK:              "成功",
		CODE_SYSTEM_ERROR:    "系统错误",
		CODE_PARAM_WRONG:     "参数错误",
		CODE_NOT_LOGINED:     "未登录",
		CODE_ILLEGAL_REQUEST: "非法请求",
		CODE_NO_PERMISSION:   "没有权限",
		CODE_NOT_FOUND:       "资源未找到",
	}
)
