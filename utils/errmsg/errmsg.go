package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块
	ERROR_USERNAME_USED      = 601
	ERROR_PASSWORD_WRONG     = 602
	ERROR_USERNAME_NOT_EXIST = 603
)

var errorMsg = map[int]string{
	SUCCESS:                  "ok",
	ERROR:                    "出现错误",
	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "输入密码错误",
	ERROR_USERNAME_NOT_EXIST: "该用户不存在",
}

func GetErrMsg(code int) string {
	return errorMsg[code]
}
