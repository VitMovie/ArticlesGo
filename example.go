package goarticles

import (
	"fmt"
)

func Example() {

	// Create article
	a := &Article{Title: "go article", Endpoint: "/api/go_article"}
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

