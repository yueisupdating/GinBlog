package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块
	ERROR_USERNAME_USED          = 601
	ERROR_PASSWORD_WRONG         = 602
	ERROR_USER_NOT_EXIST         = 603
	ERROR_TOKEN_NOT_EXIST        = 604
	ERROR_TOKEN_RUNTIME          = 605
	ERROR_TOKEN_WRONG            = 606
	ERROR_TOKEN_TYPE_WRONG       = 607
	ERROR_USER_RIGHT_WRONG       = 608
	ERROR_USER_PROFILE_NOT_EXIST = 609

	// 分类模块
	ERROR_CATENAME_USED  = 701
	ERROR_CATE_NOT_EXIST = 702
	ERROR_CATE_REPEAT    = 703

	// 文章模块
	ERROR_Article_NOT_EXIST = 801
)

var errorMsg = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "出现错误",
	ERROR_USERNAME_USED:          "用户名已存在",
	ERROR_PASSWORD_WRONG:         "输入密码错误",
	ERROR_USER_NOT_EXIST:         "该用户不存在",
	ERROR_CATENAME_USED:          "分类名已存在",
	ERROR_CATE_NOT_EXIST:         "该分类名不存在",
	ERROR_Article_NOT_EXIST:      "文章不存在",
	ERROR_TOKEN_NOT_EXIST:        "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:          "TOKEN已过期",
	ERROR_TOKEN_WRONG:            "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:       "TOKEN格式错误",
	ERROR_USER_RIGHT_WRONG:       "该用户不存在该权限",
	ERROR_CATE_REPEAT:            "分类名未修改",
	ERROR_USER_PROFILE_NOT_EXIST: "该用户个人信息不存在",
}

func GetErrMsg(code int) string {
	return errorMsg[code]
}
