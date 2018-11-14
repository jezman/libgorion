package libgorion

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestWorker(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"firstName", "midName", "lastName", "Company"}
	rows := sqlmock.NewRows(column).
		AddRow("f1", "m1", "l1", "c1").
		AddRow("f2", "m2", "l2", "c2")

	mock.ExpectQuery(testQueryEmployees).WillReturnRows(rows)

	if _, err = lib.Workers(""); err != nil {
		t.Errorf("error was not expected while gets worker %q ", err)
	}

	mock.ExpectQuery(testQueryEmployeesByCompany).WithArgs("company").WillReturnRows(rows)

	if _, err = lib.Workers("company"); err != nil {
		t.Errorf("error was not expected while gets worker by company %q ", err)
	}
}
