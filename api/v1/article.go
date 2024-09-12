package v1

import (
	"context"

	"ginblog/model"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetViewCount(c *gin.Context) {
	articleId := c.Param("id")
	count, err := utils.RedisClient.ZScore(context.Background(), "view", articleId).Result()
	if err != nil {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    strconv.Itoa(int(count)),
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArtRank(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	var rangeLeft int64 = int64((pageNum-1)*pageSize + 1)
	var rangeRight int64 = int64(pageNum * pageSize)

	members, err := utils.RedisClient.ZRange(context.Background(), "view", rangeLeft, rangeRight).Result()

	if err != nil {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    nil,
			"total":   0,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	data, code, total := model.GetArtRank(members, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func IncrViewCount(c *gin.Context) {
	articleId := c.Param("id")
	err := utils.RedisClient.ZIncrBy(context.Background(), "view", 1, articleId).Err()
	if err != nil {
		code := errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	GetViewCount(c)
}

func AddArt(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	code := model.CreateArticle(&article)
	if code == errmsg.SUCCESS {
		utils.RedisClient.ZAdd(context.Background(), "view", redis.Z{Score: 0, Member: article.ID})
	}
	ctx.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    article,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateArticle(ctx *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	_ = ctx.ShouldBindJSON(&article)

	code := model.EditArt(id, &article)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticles(ctx *gin.Context) {
	title := ctx.Query("title")
	articles, total := model.GetArticles(title)
	code := errmsg.SUCCESS

	for i, _ := range articles {
		count, _ := utils.RedisClient.ZScore(context.Background(), "view", strconv.Itoa(int(articles[i].ID))).Result()
		articles[i].ViewCount = int(count)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    articles,
		"total":   total,
	})
}

func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetArt(title, pageSize, pageNum)

	for i, _ := range data {
		count, _ := utils.RedisClient.ZScore(context.Background(), "view", strconv.Itoa(int(data[i].ID))).Result()
		data[i].ViewCount = int(count)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticleByCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, code, total := model.GetArticleByCategory(id, pageSize, pageNum)

	for i, _ := range data {
		count, _ := utils.RedisClient.ZScore(context.Background(), "view", strconv.Itoa(int(data[i].ID))).Result()
		data[i].ViewCount = int(count)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticleByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, code := model.GetArticleByID(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    user,
	})
}
