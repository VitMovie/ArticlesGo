package goarticles

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db = GetDB()
	db.AutoMigrate(&Article{})
}

// Article ...
type Article struct {
	gorm.Model
	Title    string
	ConfluenceID int
}

func (a *Article) CreateArticle() (*Article, error) {
	db.NewRecord(a)
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Article) GetAllArticles() ([]Article, error) {
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) GetArticleById(id int64) (*Article, error) {
	var article Article
	if err := db.Where("ID = ?", id).Find(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *Article) DeleteArticle(id int64) Article {
	var article Article
	db.Where("ID = ?", id).Delete(article)
	return article
}

func (a *Article) GetArticleByConfluenceId(id int64) (*Article, error) {
	var article Article
	if err := db.Where("ConfluenceID = ?", id).Find(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
