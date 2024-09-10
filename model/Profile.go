package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(20)" json:"name" label:"用户名"`
	Img         string `gorm:"column:img;type:varchar(200)" json:"img" label:"用户背景图"`
	Avatar      string `gorm:"column:avatar;type:varchar(200)" json:"avatar" label:"用户头像"`
	Description string `gorm:"column:description;type:varchar(100)" json:"description" label:"个人介绍"`
	Qq          string `gorm:"column:qq;type:varchar(30)" json:"qq" label:"qq号码"`
	WeChat      string `gorm:"column:weChat;type:varchar(30)" json:"weChat" label:"微信号"`
	Email       string `gorm:"column:email;type:varchar(30)" json:"email" label:"电子邮箱地址"`
}

func GetProfile(id int) (*Profile, int) {
	var profile *Profile
	err := db.Where("id=?", id).Find(&profile).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

func UpdateProfile(data *Profile, id int) (*Profile, int) {
	var profile *Profile
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	maps["img"] = data.Img
	maps["qq"] = data.Qq
	maps["wechat"] = data.WeChat
	maps["email"] = data.Email
	maps["description"] = data.Description
	maps["avatar"] = data.Avatar

	err := db.Model(profile).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}
