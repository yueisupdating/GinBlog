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
	UserName string `gorm:"column:username;type:varchar(20);not null " json:"username"`
	Password string `gorm:"column:password;type:varchar(200);not null " json:"password"`
	Role     int    `gorm:"column:role;type:int;not null " json:"role"`
}

func CreateUser(user *User) int {
	user.Password = ScryptPassword(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckUserName(name string) int {
	var user User
	db.Select("id").Where("username=?", name).First(&user)

	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := db.Delete("id").Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//TODO: 编辑用户

func GetUsers(pageSize int, pageNum int) ([]User, int) {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return users, errmsg.SUCCESS
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
