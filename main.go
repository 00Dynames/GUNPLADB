package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/00Dynames/GUNPLADB/resources"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var dbConn *sql.DB

type Env struct {
	DB resources.DB
}

func main() {

	db, err := resources.OpenConnection()
	defer resources.CloseConnection(db)

	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/1.0/gunpla/grades", func(w http.ResponseWriter, r *http.Request) {
		grades, _ := db.GetGrades()
		message, _ := json.Marshal(grades)
		w.Write(message)
	}).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/grades/{grade_id}", func(w http.ResponseWriter, r *http.Request) {
		grade := mux.Vars(r)["grade_id"]
		kits, _ := db.GetGradeKits(&grade, nil)
		message, _ := json.Marshal(kits)
		w.Write(message)
	}).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/grades/{grade_id}/{kit_id}", func(w http.ResponseWriter, r *http.Request) {
		grade := mux.Vars(r)["grade_id"]
		kit, _ := strconv.Atoi(mux.Vars(r)["kit_id"])
		result, _ := db.GetGradeKits(&grade, &kit)
		message, _ := json.Marshal(result[0])
		w.Write(message)
	}).Methods("GET")
	//r.HandleFunc("/api/1.0/gunpla/series/{series_id}", series_get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
