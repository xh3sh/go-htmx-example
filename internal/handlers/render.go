package handlers

import (
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	funcMap := template.FuncMap{
		"in": func(slice []string, item string) bool {
			for _, v := range slice {
				if v == item {
					return true
				}
			}
			return false
		},
		"join": strings.Join,
		"toggleTag": func(slice []string, item string) []string {
			for i, v := range slice {
				if v == item {
					return append(slice[:i], slice[i+1:]...)
				}
			}
			return append(slice, item)
		},
		"escape": func(s string) string {
			return strings.ReplaceAll(strings.ReplaceAll(s, `\`, `\\`), `"`, `\"`)
		},
	}
	return &Templates{
		templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("views/*.html")),
	}
}
