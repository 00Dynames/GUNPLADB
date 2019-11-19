package resources

import (
	"database/sql"
	"fmt"
	"log"
)

type Grade struct {
	Id   int
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (db *DB) GetGrades() ([]Grade, error) {
	result := []Grade{}
	qResult, err := db.Query("select * from grades")
	if err != nil {
		log.Panic(err)
	}

	if err != nil {
		return nil, err
	}

	for qResult.Next() {
		g := Grade{}
		qResult.Scan(&g.Id, &g.Name)
		//TODO: replace localhost with a parameterised base url
		//TODO: replace name with id in url after enabling the endpoint
		//			to take name or id
		g.Url = fmt.Sprintf("%s/api/1.0/gunpla/grades/%v", "localhost:8080", g.Name)
		result = append(result, g)
	}

	return result, nil
}

func (db *DB) GetGradeKits(grade *string, kit *int) ([]Kit, error) {
	result := []Kit{}
	var err error

	//fmt.Printf("%s, %d", *grade, *kit)
	//TODO: do something if the grade id is out of range

	var tmp *sql.Rows
	if kit != nil {
		tmp, err = db.Query(
			fmt.Sprintf(
				"select * from gunpla where grade='%s' and grade_id='%d'",
				*grade,
				*kit,
			),
		)
	} else {
		tmp, err = db.Query(
			fmt.Sprintf(
				"select * from gunpla where grade='%s'",
				*grade,
			),
		)
	}

	if err != nil {
		return nil, err
	}

	for tmp.Next() {
		k := Kit{}
		err = tmp.Scan(
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
