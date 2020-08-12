package main

import (
	"fmt"
	"github.com/vitmovie/articlesgo"
)

const (
	username = "root"
	password  = "root"
	host = "localhost"
	dbname = "confluence"
)

func main() {
	// Connect into db
	db := goarticles.NewDatabase(username, password, host, dbname)
	// Migrate
	db.Connection.AutoMigrate(&goarticles.Article{})

	// Create article
	a := &goarticles.Article{Title: "go article", ConfluenceID: 22, DB: db}
	_, err := a.CreateArticle()
	if err != nil {
		panic(err)
	}
	id := a.ID

	a = &goarticles.Article{DB: db}

	// Get Article
	article, err := a.GetArticleById(int64(id))
	if err != nil {
		panic(err)
	}
	fmt.Println(article)

	// Get Articles
	articles, err := a.GetAllArticles()
	if err != nil {
		panic(err)
	}
	fmt.Println(articles)

	// Delete Article
	a.DeleteArticle(int64(id))
}

