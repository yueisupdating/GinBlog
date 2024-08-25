package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	// 继承model.go, 包括时间戳、ID作主键
	gorm.Model
	UserName string `gorm:"column:username;type:varchar(20);not null " json:"username" validate:"required,min=4,max=20" label:"用户名"`
	Password string `gorm:"column:password;type:varchar(200);not null " json:"password" validate:"required,min=5,max=50" label:"密码"`
	Role     int    `gorm:"column:role;type:int;DEFAULT:2 " json:"role" validate:"required,gte=2" label:"角色"`
}

// 创建新用户
func CreateUser(user *User) int {
	user.Password = ScryptPassword(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查看用户名是否已存在
func CheckUserName(name string) int {
	var user User
	db.Select("id").Where("username=?", name).First(&user)

	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

/*
如果你的模型包含了 gorm.DeletedAt字段（该字段也被包含在gorm.Model中），
那么该模型将会自动获得软删除的能力！
当调用Delete时，GORM并不会从数据库中删除该记录，而是将该记录的DeleteAt设置为当前时间，
而后的一般查询方法将无法查找到此条记录。
*/

// 判断用户是否存在
func IsUserExist(id int) int {
	var user User
	db.Where("id=?", id).Find(&user)
	if user.ID > 0 {
		return 1
	}
	return 0
}

// 软删除用户
func DeleteUser(id int) int {
	var user User
	if IsUserExist(id) == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	if IsUserExist(id) == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	maps := make(map[string]interface{})
	maps["username"] = data.UserName
	maps["role"] = data.Role
	err := db.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查找用户
func GetUser(id int) (User, int) {
	var user User
	db.Where("id=?", id).First(&user)
	if user.ID <= 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}

// 返回当前所有用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// Scrypt密码加密
func ScryptPassword(password string) string {
	salt := []byte{5, 21, 88, 231, 114, 199, 31, 67}
	scryptPw, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Panicln(err)
	}
	return base64.StdEncoding.EncodeToString(scryptPw)
}

func CheckLogin(username string, password string) (User, int) {
	var user User
	db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_RIGHT_WRONG
	}
	if ScryptPassword(password) != user.Password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}
