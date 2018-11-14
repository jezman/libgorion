package libgorion

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Company", "Workers"}).
		AddRow("company 1", "2").
		AddRow("company 2", "5")

	mock.ExpectQuery(testQueryCompanies).WillReturnRows(rows)

	lib := &Database{DB: db}
	if _, err = lib.Company(); err != nil {
		t.Errorf("error was not expected while gets company %q ", err)
	}
}
