package resources

import (
	"database/sql"
	"fmt"
	"log"
)

type grade struct {
	Id    int
	Grade string
}

func (db *DB) GetGrades() ([]grade, error) {
	result := []grade{}
	qResult, err := db.Query("select * from grades")
	if err != nil {
		log.Panic(err)
	}

	if err != nil {
		return nil, err
	}

	for qResult.Next() {
		g := grade{}
		qResult.Scan(&g.Id, &g.Grade)
		result = append(result, g)
	}

	return result, nil
}

func (db *DB) GetGradeKits(grade *string, kit *int) ([]Kit, error) {
	result := []Kit{}
	var err error

	//TODO: do something if the grade id is out of range

	var tmp *sql.Rows
	if kit != nil {
		tmp, err = db.Query(
			fmt.Sprintf(
				"select * from gunpla where grade='%s' and grade_id='%s'",
				grade,
				kit,
			),
		)
	} else {
		tmp, err = db.Query(
			fmt.Sprintf(
				"select * from gunpla where grade='%s'",
				grade,
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
