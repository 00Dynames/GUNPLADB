package resources

import (
	"fmt"
)

type Series struct {
	Id   int
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (db *DB) GetSeries() ([]Series, error) {
	result := []Series{}
	qResult, err := db.Query("select * from series")
	if err != nil {
		return nil, err
	}

	for qResult.Next() {
		s := Series{}
		qResult.Scan(&s.Id, &s.Name)
		s.Url = fmt.Sprintf("%s/api/1.0/gunpla/series/%v", "localhost:8080", s.Id)
		result = append(result, s)
	}

	return result, nil
}

/*func series_get(w http.ResponseWriter, r *http.Request) {

	result, err := dbQuery(
		dbConn,
		fmt.Sprintf("select * from gunpla where series = %s", mux.Vars(r)["series"]),
	)
	logError(err)

	message, err := json.Marshal(result)
	logError(err)
	w.Write(message)
}*/
