package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orange432/bloggo/pkg/articles"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Bloggo.",
	})
}

func About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{
		"Title": "Bloggo. - About",
	})
}

func Article(c *fiber.Ctx) error {
	data := articles.FakeArticle()
	return c.Render("article", fiber.Map{
		"articleTitle":   data.Title,
		"articleContent": data.Content,
		"articleAuthor":  data.Author,
	})
}
