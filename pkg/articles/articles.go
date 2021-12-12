package articles

import (
	"fmt"
	"io/ioutil"

	"github.com/orange432/bloggo/pkg/config"
	"github.com/orange432/bloggo/pkg/models"
)

var app *config.AppConfig

func InitializeArticles(a *config.AppConfig) {
	app = a
}

// func LoadFSArticle(slug string) models.Article {
// 	data, err := LoadFileContent("./articles/" + slug + ".txt")
// 	if err != nil {
// 		log.Println(err)
// 		return models.Article{Title: "Error!", Content: "Error, something went wrong!", CachedAt: 0}
// 	}
// 	ac := string(data[:])
// 	return models.Article{
// 		Title:    slug,
// 		Content:  ac,
// 		CachedAt: int(time.Now().Unix()),
// 	}
// }

func LoadFileContent(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("error loading file", filepath)
		return nil, err
	}
	return data, nil
}

func SaveArticle(slug string, article models.Article) error {
	return nil
}
