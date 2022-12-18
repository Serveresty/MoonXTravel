package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type GetData struct {
	from      string
	to        string
	when_from string
	when_to   string
	people    string
}

type SendData struct {
	from      string
	to        string
	when_from string
	when_to   string
	people    string
}

type DataBase struct {
	Data *sql.DB
}

func (s *DataBase) SearchTickets(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var get_data GetData

	err := json.NewDecoder(r.Body).Decode(&get_data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var send_data SendData

}
