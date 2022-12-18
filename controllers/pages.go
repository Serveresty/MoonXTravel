package controllers

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var files []string

func MainPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files = []string{
		"./static/templates/main_page.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}

func SearchResultsPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files = []string{
		"./static/templates/search_result.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}
