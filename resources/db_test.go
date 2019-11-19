package resources_test

import (
	"reflect"
	"testing"

	"github.com/00Dynames/GUNPLADB/resources"
	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {

	//mockCtrl := gomock.NewController(t)
	//defer mockCtrl.Finish()

	db, err := resources.OpenConnection()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if reflect.TypeOf(db) != reflect.TypeOf((*resources.DB)(nil)) {
		t.Errorf(
			"Expected OpenConnectiont to return %v, got %v",
			reflect.TypeOf((*resources.DB)(nil)),
			reflect.TypeOf(db),
		)
	}

}
