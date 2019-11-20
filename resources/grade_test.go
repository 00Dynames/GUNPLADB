package resources_test

import (
	"fmt"
	"testing"

	"github.com/00Dynames/GUNPLADB/resources"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetGrades(t *testing.T) {
	mockDB, m, _ := sqlmock.New()
	db := resources.DB{mockDB}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "RG").AddRow(2, "MG")
	m.ExpectQuery("select \\* from grades").WillReturnRows(rows)

	expected := []resources.Grade{
		{1, "RG", "localhost:8080/api/1.0/gunpla/grades/RG"},
		{2, "MG", "localhost:8080/api/1.0/gunpla/grades/MG"},
	}

	result, err := db.GetGrades()
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestGetGradesQueryFail(t *testing.T) {

	mockDB, m, _ := sqlmock.New()
	db := resources.DB{mockDB}

	// TODO: it *might* matter what the error is, maybe?
	m.ExpectQuery("select \\* from grades").WillReturnError(fmt.Errorf("some error"))

	_, err := db.GetGrades()
	assert.Equal(t, fmt.Errorf("some error"), err)
}

func TestGetGradeKits(t *testing.T) {

	mockDB, m, _ := sqlmock.New()
	db := resources.DB{mockDB}

	rows := sqlmock.NewRows(
		[]string{
			"id",
			"grade_id",
			"grade",
			"name",
			"series",
			"price",
			"release",
			"description",
		},
	).AddRow(1, 1, "RG", "RX-78-2", "Mobile Suit Gundam", "1000", "", "Test")
	m.ExpectQuery("select \\* from gunpla where grade='RG'").WillReturnRows(rows)

	expected := []resources.Kit{{1, 1, "RG", "RX-78-2", "Mobile Suit Gundam", 1000, "", "Test"}}
	grade := "RG"
	result, err := db.GetGradeKits(&grade, nil)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}
