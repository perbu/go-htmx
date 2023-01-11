package tmpls

import (
	_ "embed"
	"html/template"
)

//go:embed poem.gotmpl
var poem string

type Templates struct {
	templates map[string]*template.Template
}

func Load() (*Templates, error) {
	t := &Templates{
		templates: make(map[string]*template.Template),
	}
	t.templates["poem"] = template.Must(template.New("poem").Parse(poem))
	return t, nil
}
