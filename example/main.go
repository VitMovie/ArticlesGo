package main

import (
	"fmt"
	articlesgo "github.com/vitmovie/articlesgo"
)

const (
	username = "root"
	password  = "password"
	host = "localhost"
	dbname = "confluence"
)

func main() {
	// Connect into db
	db := articlesgo.NewDatabase(username, password, host, dbname)
	// Migrate
	db.Connection.AutoMigrate(&articlesgo.Article{})

	// Create article
	a := &articlesgo.Article{Title: "go article", ConfluenceID: 22, Url: "http://conf.net", DB: db}
	_, err := a.CreateArticle()
	if err != nil {
		panic(err)
	}
	id := a.ID

	a = &articlesgo.Article{DB: db}

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

