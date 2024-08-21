package model

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignkey:Cid"`
	Title       string   `gorm:"column:title;type:varchar(20);not null" json:"title"`
	Description string   `gorm:"column:description;type:varchar(200)" json:"description"`
	Content     string   `gorm:"column:content;type:longtext" json:"content"`
	Cid         int      `gorm:"column:cid;type:int" json:"cid"`
	Img         string   `gorm:"column:img;type:varchar(200)" json:"img"`
}

// 创建新文章
func CreateArticle(art *Article) int {
	err := db.Create(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 判断文章是否存在
func IsArticleExist(id int) int {
	var art Article
	db.Where("id=?", id).Find(&art)
	if art.ID > 0 {
		return 1
	}
	return 0
}

// 软删除文章
func DeleteArt(id int) int {
	var art Article
	if IsArticleExist(id) == 0 {
		return errmsg.ERROR_Article_NOT_EXIST
	}
	err := db.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var article Article
	if IsArticleExist(id) == 0 {
		return errmsg.ERROR_Article_NOT_EXIST
	}
	maps := make(map[string]interface{})
	maps["description"] = data.Description
	maps["img"] = data.Img
	maps["content"] = data.Content
	maps["cid"] = data.Cid
	maps["title"] = data.Title

	err := db.Model(&article).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 返回当前所有文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}
