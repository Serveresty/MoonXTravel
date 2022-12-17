package main

import (
	"MoonXTravel/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	routes(r)

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/static/*filepath", http.Dir("static"))

	r.GET("/", controllers.MainPage)
}
