package models

import "html/template"

type TemplateData struct {
	StringMap map[string]string
	Article   template.HTML
	Success   bool
	Error     string
}
