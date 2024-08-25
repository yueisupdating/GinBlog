package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint   `gorm:"column:id;primary_key;auto_increment" json:"id"`
	CategoryName string `gorm:"column:categoryname;type:varchar(20);not null" json:"categoryname"`
}

// 创建新分类
func CreateCategory(cate *Category) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	if IsCateExist(id) == 0 {
		return errmsg.ERROR_CATE_NOT_EXIST
	}
	err := db.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑分类
func EditCate(id int, data *Category) int {
	if IsCateExist(id) == 0 {
		return errmsg.ERROR_CATE_NOT_EXIST
	}
	var cate Category
	maps := make(map[string]interface{})
	maps["categoryname"] = data.CategoryName
	err := db.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个分类
func GetCate(id int) (Category, int) {
	var cate Category
	db.Where("id=?", id).First(&cate)
	if cate.ID <= 0 {
		return cate, errmsg.ERROR_CATE_NOT_EXIST
	}
	return cate, errmsg.SUCCESS
}

// 返回当前所有分类
func GetCates(pageSize int, pageNum int) ([]Category, int64) {
	var cates []Category
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// 查看分类名是否已存在
func CheckCateName(name string) int {
	var cate Category
	db.Select("id").Where("categoryname=?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 判断分类是否存在
func IsCateExist(id int) int {
	var cate Category
	db.Where("id=?", id).Find(&cate)
	if cate.ID > 0 {
		return 1
	}
	return 0
}
