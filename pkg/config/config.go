package config

import (
	"html/template"

	"github.com/orange432/bloggo/pkg/models"
)

type AppConfig struct {
	UseCache          bool
	TemplateCache     map[string]*template.Template
	UseArticleCache   bool // Enable/Disable article caching
	ArticleCacheLimit int  // 0 means there is no limit
	ArticleCache      map[string]*models.Article
}
