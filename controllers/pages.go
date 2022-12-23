package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var files []string

type DataBase struct {
	Data *sql.DB
}

type SendData struct {
	from_where  string
	to_where    string
	when_from   string
	when_to     string
	count_place string
	price       string
	earth       bool
}

func MainPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files = []string{
		"./static/templates/main_page.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}

func ResPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files = []string{
		"./static/templates/search_result.tmpl",
	}
	var tpl = template.Must(template.ParseFiles(files...))
	tpl.Execute(rw, nil)
}

func (s *DataBase) SearchResultsPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	from := r.FormValue("first-input")
	to := r.FormValue("destination")
	when_from := r.FormValue("whento")
	when_to := r.FormValue("whenback")
	cnt_passengers := r.FormValue("last-input")

	fmt.Println("Entered DATA: " + from + " " + to + " " + when_from + " " + when_to + " " + cnt_passengers)

	rows, err := s.Data.Query(`SELECT frrom, tto, when_from, when_to, count_places, price FROM ticket_to WHERE when_from=?`, when_from)
	if err != nil {
		panic(err)
	}
	var otkyda, kyda, kogda, prilet, col_vo_mest, price string
	for rows.Next() {
		err = rows.Scan(&otkyda, &kyda, &kogda, &prilet, &col_vo_mest, &price)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("--------------------------------------")
	fmt.Println("|Билет из пункта отправения 'Земля'")
	fmt.Println("|Откуда: " + otkyda + " Куда: " + kyda + " Когда: " + kogda + " Время прилёта: " + prilet + " Цена: " + price + " Кол-во билетов: " + cnt_passengers)
	fmt.Println("--------------------------------------")
	fmt.Println("")

	rows1, err := s.Data.Query(`SELECT frrom, tto, when_from, when_to, count_places, price FROM ticket_to WHERE when_from=?`, when_to)
	if err != nil {
		panic(err)
	}
	var otkyda1, kyda1, kogda1, prilet1, col_vo_mest1, price1 string
	for rows1.Next() {
		err = rows1.Scan(&otkyda1, &kyda1, &kogda1, &prilet1, &col_vo_mest1, &price1)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("--------------------------------------")
	fmt.Printf("|Билет из пункта отправения '%s'", otkyda1)
	fmt.Println("")
	fmt.Println("|Откуда: " + otkyda1 + " Куда: " + kyda1 + " Когда: " + kogda1 + " Время прилёта: " + prilet1 + " Цена: " + price1 + " Кол-во билетов: " + cnt_passengers)
	fmt.Println("--------------------------------------")
	fmt.Println("")
	http.Redirect(rw, r, "/results", http.StatusSeeOther)
	return
}
