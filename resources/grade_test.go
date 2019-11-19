package resources_test

import (
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

	result, _ := db.GetGrades()
	assert.Equal(t, expected, result)
}
