package goarticles

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Article ...
type Article struct {
	gorm.Model
	Title        string
	ConfluenceID int
	Url 		 string
	DB           *Database `gorm:"-"`
}

func (a *Article) CreateArticle() (*Article, error) {
	a.DB.Connection.NewRecord(a)
	if err := a.DB.Connection.Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Article) GetAllArticles() ([]Article, error) {
	var articles []Article
	if err := a.DB.Connection.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) GetArticleById(id int64) (*Article, error) {
	var article Article
	if err := a.DB.Connection.Where("ID = ?", id).Find(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *Article) DeleteArticle(id int64) Article {
	var article Article
	a.DB.Connection.Where("ID = ?", id).Delete(article)
	return article
}

func (a *Article) GetArticleByConfluenceId(id int64) (*Article, error) {
	var article Article
	if err := a.DB.Connection.Where("ConfluenceID = ?", id).Find(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
