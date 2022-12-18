package main

import (
	"MoonXTravel/controllers"
	"MoonXTravel/service"
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	DB, errdb := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tickets_data")
	if errdb != nil {
		panic(errdb)
	}
	defer DB.Close()
	if err := DB.Ping(); err != nil {
		panic(err)
	}

	r := httprouter.New()
	routes(r, DB)

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router, DB *sql.DB) {
	r.ServeFiles("/static/*filepath", http.Dir("static"))

	handler := service.DataBase{Data: DB}

	r.GET("/", controllers.MainPage)
	r.GET("/results", controllers.SearchResultsPage)
	r.POST("/results", handler.SearchTickets)
}
