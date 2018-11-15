package libgorion

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestDoors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"GIndxex", "Name"}).
		AddRow(1, "door 1").
		AddRow(2, "door 2")

	mock.ExpectQuery("^SELECT (.+), (.+) FROM AcessPoint ORDER BY GIndex$").WillReturnRows(rows)

	lib := &Database{DB: db}
	if _, err = lib.Doors(); err != nil {
		t.Errorf("error was not expected while gets doors %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
