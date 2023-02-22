package errmsg

const (
	UserUsed       = 1001
	UserNotExist   = 1002
	UserLoginWrong = 1003

	TokenNotExist    = 2001
	TokenExpired     = 2002
	TokenNotValidYet = 2003
	TokenMalformed   = 2004
	TokenInvalid     = 2005
)

var MsgMap = map[int]string{
	UserUsed:       "用户名已被注册",
	UserNotExist:   "用户不存在",
	UserLoginWrong: "账号或密码错误",

	TokenNotExist:    "授权token不存在",
	TokenExpired:     "token已过期，请重新登录",
	TokenNotValidYet: "token无效",
	TokenMalformed:   "token不正确",
	TokenInvalid:     "token非法",
}

func GetErrMsg(code int) string {
	return MsgMap[code]
}
