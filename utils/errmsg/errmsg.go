package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块
	ERROR_USERNAME_USED  = 601
	ERROR_PASSWORD_WRONG = 602
	ERROR_USER_NOT_EXIST = 603
	// 分类模块
	ERROR_CATENAME_USED  = 701
	ERROR_CATE_NOT_EXIST = 702
	// 文章模块
	ERROR_Article_NOT_EXIST = 801
)

var errorMsg = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "出现错误",
	ERROR_USERNAME_USED:     "用户名已存在",
	ERROR_PASSWORD_WRONG:    "输入密码错误",
	ERROR_USER_NOT_EXIST:    "该用户不存在",
	ERROR_CATENAME_USED:     "分类名已存在",
	ERROR_CATE_NOT_EXIST:    "该分类名不存在",
	ERROR_Article_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return errorMsg[code]
}
