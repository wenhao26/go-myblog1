package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"blog/api"
	"blog/models"
)

type Article struct {
	*api.Base
}

var ArticleAPI = &Article{
	Base: api.BaseAPI,
}

// 获取文章列表
func (a *Article) Articles(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	list, total := models.GetArticles(pageInt, limitInt)
	a.Success(ctx, "success", map[string]interface{}{
		"list":  list,
		"total": total,
	})
}

// 获取文章信息
func (a *Article) Article(ctx *gin.Context) {
	aid := ctx.Query("aid")
	aidInt, _ := strconv.Atoi(aid)

	info := models.GetArticle(aidInt)
	a.Success(ctx, "success", info)
}
