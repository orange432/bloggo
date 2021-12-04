package articles

import (
	"io/ioutil"
	"log"
)

type Article struct {
	Slug      string
	Title     string
	Content   string
	Author    string
	Published int
}

// ListArticles creates a list of all available articles
func ListArticles() (map[string]Article, error) {
	articleMap := make(map[string]Article)

	files, err := ioutil.ReadDir("./posts/")
	if err != nil {
		log.Println("Unable to load articles directory!")
		return articleMap, err
	}

	for _, file := range files {
		articleMap[file.Name()] = Article{
			Slug:      file.Name(),
			Title:     "TEST MAP",
			Content:   "test",
			Author:    "orange432",
			Published: 14040,
		}
	}

	return articleMap, nil
}

func LoadArticle(slug string) *Article {
	arti := Article{
		Slug:      slug,
		Title:     "test",
		Content:   "test",
		Author:    "orange432",
		Published: 14040,
	}
	return &arti
}

//FakeArticle generates a fake article for testing
func FakeArticle() *Article {
	arti := Article{
		Slug:      "test-slug",
		Title:     "test",
		Content:   "test",
		Author:    "orange432",
		Published: 14040,
	}
	return &arti
}
