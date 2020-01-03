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

func (db *DB) GetSeriesKits(series *int) ([]Kit, error) {

	result := []Kit{}

	qResult, err := db.Query(fmt.Sprintf("select * from gunpla where series =(select series from series where id = %d)", *series))
	if err != nil {
		return nil, err
	}

	for qResult.Next() {
		k := Kit{}
		err := qResult.Scan(
			&k.Id,
			&k.Grade_id,
			&k.Grade,
			&k.Name,
			&k.Series,
			&k.Price,
			&k.Release,
			&k.Description,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, k)
	}

	return result, nil
}
