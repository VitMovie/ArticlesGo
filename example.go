package goarticles

import (
	"fmt"
)

const (
	username = "root"
	password  = "root"
	host = "localhost"
	dbname = "confluence"
)

func Example() {
	// Connect into db
	Connect(username, password, host, dbname)

	// Create article
	a := &Article{Title: "go article", ConfluenceID: 22}
	_, err := a.CreateArticle()
	if err != nil {
		panic(err)
	}
	id := a.ID

	a = &Article{}

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

