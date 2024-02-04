package templates

import (
	tlib "html/template"
	"net/http"
)

type Template struct {
	Tmpl *tlib.Template
}

type Templater interface {
	ExecuteTemplate(w http.ResponseWriter, templateName string, data interface{}) error
}

func (t *Template) LoadTemplates(pattern string) error {
	tmpl, err := tlib.ParseGlob(pattern)
	if err != nil {
		return err
	}

	t.Tmpl = tmpl
	return nil
}

func (t *Template) ExecuteTemplate(w http.ResponseWriter, templateName string, data interface{}) error {
	err := t.Tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		return err
	}

	return nil
}
