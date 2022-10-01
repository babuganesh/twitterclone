package views

import (
	"html/template"
	"net/http"
)

func NewView(layout string, files ...string) *View {
	files = append(files,
		// "views/navbar.gohtml",
		// "views/home.gohtml",
		"views/layouts/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
