package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

type grade struct {
	Id    int
	Grade string
}

var dbConn *sql.DB

func logError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func dbQuery(conn *sql.DB, query string) ([]gunpla_kit, error) {

	result := []gunpla_kit{}

	fmt.Println(query)

	qResult, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for qResult.Next() {
		var kit gunpla_kit
		err := qResult.Scan(
			&kit.Id,
			&kit.Grade_id,
			&kit.Grade,
			&kit.Name,
			&kit.Series,
			&kit.Price,
			&kit.Release,
			&kit.Description,
		)

		result = append(result, kit)

		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func getGrades(w http.ResponseWriter, r *http.Request) {

	result := []grade{}
	qResult, _ := dbConn.Query("select * from grades")
	fmt.Println(result)

	for qResult.Next() {
		g := grade{}
		qResult.Scan(&g.Id, &g.Grade)
		result = append(result, g)
	}

	message, _ := json.Marshal(result)
	w.Write(message)

}

func gunpla_get(w http.ResponseWriter, r *http.Request) {
	result := []gunpla_kit{}
	var err error
	grade := mux.Vars(r)["grade_id"]
	kit := ""

	//TODO: do something if the grade id is out of range
	if mux.Vars(r)["kit_id"] != "" {
		kit = mux.Vars(r)["kit_id"]
	}

	if kit != "" {
		result, err = dbQuery(
			dbConn,
			fmt.Sprintf(
				"select * from gunpla where grade='%s' and grade_id='%s'",
				grade,
				kit,
			),
		)
	} else {
		result, err = dbQuery(
			dbConn,
			fmt.Sprintf(
				"select * from gunpla where grade='%s'",
				grade,
			),
		)
	}
	logError(err)

	message, err := json.Marshal(result)
	logError(err)
	w.Write(message)
}

func series_get(w http.ResponseWriter, r *http.Request) {

	result, err := dbQuery(
		dbConn,
		fmt.Sprintf("select * from gunpla where series = %s", mux.Vars(r)["series"]),
	)
	logError(err)

	message, err := json.Marshal(result)
	logError(err)
	w.Write(message)
}

func main() {

	var err error
	dbConn, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/gunpladb",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
		),
	)
	logError(err)

	defer dbConn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/api/1.0/gunpla/grades", getGrades).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/grades/{grade_id}", gunpla_get).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/grades/{grade_id}/{kit_id}", gunpla_get).Methods("GET")
	r.HandleFunc("/api/1.0/gunpla/series/{series_id}", series_get).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
