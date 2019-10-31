package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type gunpla_kit struct {
	Id          int    `json:"id"`
	Grade_id    int    `json:"grade_id"`
	Grade       string `json:"grade"`
	Name        string `json:"name"`
	Series      string `json:"series"`
	Price       int    `json:"price"`
	Release     string `json:"release"`
	Description string `json:"description"`
}

//TODO: wrap err checking conditions into a function

func gunpla_get(w http.ResponseWriter, r *http.Request) {
	jsonData := []gunpla_kit{}
	grade := mux.Vars(r)["grade"]
	grade_id := ""

	//TODO: do something if the grade id is out of range
	if mux.Vars(r)["grade_id"] != "" {
		grade_id = mux.Vars(r)["grade_id"]
	}

	//TODO: parameterise mysql creds
	db, err := sql.Open("mysql", "dbunadi:bcWoJwgiO81AaNDMj1oE@tcp(gunpladb-1.clqhihsn26ab.ap-southeast-2.rds.amazonaws.com)/gunpladb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var results *sql.Rows
	if grade_id != "" {
		results, err = db.Query(fmt.Sprintf("select * from gunpla where grade='%s' and grade_id='%s'", grade, grade_id))
	} else {
		results, err = db.Query(fmt.Sprintf("select * from gunpla where grade='%s'", grade))
	}

	if err != nil {
		log.Printf("query error")
		log.Fatal(err)
	}
	for results.Next() {
		var kit gunpla_kit
		err := results.Scan(
			&kit.Id,
			&kit.Grade_id,
			&kit.Grade,
			&kit.Name,
			&kit.Series,
			&kit.Price,
			&kit.Release,
			&kit.Description,
		)

		if err != nil {
			log.Fatal(err)
		}

		jsonData = append(jsonData, kit)
	}

	message, err := json.Marshal(jsonData)
	w.Write(message)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/1.0/gunpla/{grade}", gunpla_get).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/{grade}/{grade_id}", gunpla_get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
