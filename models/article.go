package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title        string
	Cid          int
	Desc         string
	Content      string
	Img          string
	CommentCount int
	ReadCount    int
}

// 获取文章列表
func GetArticles(page, limit int) ([]Article, int64) {
	var articles []Article
	var total int64

	err := db.Limit(limit).Offset((page - 1) * limit).Find(&articles).Error
	if err != nil {
		return articles, 0
	}

	db.Model(&articles).Count(&total)
	return articles, total
}

// 获取文章信息
func GetArticle(aid int) Article {
	var article Article
	db.Where("id=?", aid).First(&article)
	return article
}
