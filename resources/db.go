package resources

import (
	"database/sql"
	"fmt"
	"os"
)

type ResourceData interface {
	GetGrades() ([]grade, error)
}

type DB struct {
	*sql.DB
}

func OpenConnection() (*DB, error) {
	// TODO: throw an error if any of the env variables aren't set
	result, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/gunpladb",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
		),
	)

	if err != nil {
		return nil, err
	}

	return &DB{result}, nil
}

func CloseConnection(db *DB) {
	// TODO: check for closing connection errors
	db.Close()
}
