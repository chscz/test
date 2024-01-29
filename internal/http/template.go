package http

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// Define the template registry struct
type TemplateRegistry struct {
	Templates *template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func RegistTemplate() *TemplateRegistry {
	return &TemplateRegistry{
		Templates: template.Must(template.ParseGlob("internal/http/static/template/*.html")),
	}
}
