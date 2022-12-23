package main

import (
	"MoonXTravel/controllers"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	DB, errdb := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tickets")
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

	handler := controllers.DataBase{Data: DB}

	r.GET("/", controllers.MainPage)
	r.GET("/results", controllers.ResPage)
	r.POST("/results", handler.SearchResultsPage)
}
