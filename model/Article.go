package model

import (
	"ginblog/utils/errmsg"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignKey:Cid;references:ID"`
	Title       string   `gorm:"column:title;type:varchar(20);not null" json:"title"`
	Description string   `gorm:"column:description;type:varchar(200)" json:"description"`
	Content     string   `gorm:"column:content;type:longtext" json:"content"`
	Cid         int      `gorm:"column:cid;type:int" json:"cid"`
	Img         string   `gorm:"column:img;type:varchar(200)" json:"img"`
	ViewCount   int      `gorm:"column:viewCount;type:int;default:0" json:"viewCount"`
}

// 创建新文章
func CreateArticle(art *Article) int {
	err := db.Create(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	db.Preload("Category").First(&art, art.ID)
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

func GetArtRank(idList []string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var idLists []int

	total := len(idList)
	for i := 0; i < total; i++ {
		idTmp, _ := strconv.Atoi(idList[i])
		idLists = append(idLists, idTmp)
	}
	err = db.Preload("Category").Where("id IN ?", idLists).Find(&articleList).Error

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, int64(total)
}

func GetArt(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err := db.Preload("Category").Order("Created_At DESC").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error

	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 返回当前所有文章列表
func GetArticles(title string) ([]Article, int64) {
	var articles []Article
	var total int64

	err := db.Preload("Category").Where("title LIKE ?", title+"%").Find(&articles).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return articles, total
}

// 单个查询
func GetArticleByID(id int) (Article, int) {
	var article Article
	err := db.Where("id=?", id).Preload("Category").First(&article).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return article, errmsg.ERROR_Article_NOT_EXIST
		}
		log.Panicln(err)
		return article, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}

// 返回分类id下所有文章
func GetArticleByCategory(id int, pageSize int, pageNum int) (arts []Article, code int, cnt int64) {
	var articles []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&articles).Error
	db.Model(&articles).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return articles, errmsg.SUCCESS, total
}
