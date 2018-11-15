package libgorion

import (
	"errors"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestWorkers(t *testing.T) {
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

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindWorkerIDByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")
	mock.ExpectQuery(testQueryFindWorkerIDByName).WithArgs("Fn", "Mn", "Ln").WillReturnRows(rows)

	if err = lib.findWorkerIDByName("fn mn ln"); err != nil {
		t.Errorf("error %q ", err)
	}

	nilRows := sqlmock.NewRows(column)
	mock.ExpectQuery(testQueryFindWorkerIDByName).WithArgs("No", "Found", "Worker").WillReturnRows(nilRows)

	err = lib.findWorkerIDByName("no found worker")
	if err == errors.New("worker not found") {
		t.Errorf("error %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddWorkerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO pList").
		WithArgs("Fname", "Mname", "Lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	lib := &Database{DB: db}

	if err = lib.AddWorker("fname mname lname"); err != nil {
		t.Errorf("error was not expected while add worker: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddWorkerRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO pList").
		WithArgs("Fname", "Mname", "Lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.AddWorker("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteWorkerSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}
	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).WithArgs("Fname", "Mname", "Lname").WillReturnRows(rows)
	mock.ExpectExec(testQueryDeleteWorkerCards).
		WithArgs("fname mname lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(testQueryDeleteWorker).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = lib.DeleteWorker("fname mname lname"); err != nil {
		t.Errorf("error was not expected while delete worker: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when worker not found
func TestDeleteWorkerRollbackOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.DeleteWorker("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when worker card delete
func TestDeleteWorkerFindedRollbackTwo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryDeleteWorkerCards).
		WithArgs("fname mname lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.DeleteWorker("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when worker deleted
func TestDeleteWorkerFindedRollbackThree(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryDeleteWorkerCards).
		WithArgs("fname mname lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(testQueryDeleteWorker).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.DeleteWorker("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDisableWorkerCardSucces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryUpdateWorkerCardStatus).
		WithArgs(disableCardCode, "fname mname lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	lib := &Database{DB: db}

	if err = lib.DisableWorkerCard("fname mname lname"); err != nil {
		t.Errorf("error when disable card: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when worker not found
func TestDisableWorkerCardRollbackOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.DisableWorkerCard("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when card disable
func TestDisableWorkerCardRollbackTwo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryUpdateWorkerCardStatus).
		WithArgs(disableCardCode, "fname mname lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.DisableWorkerCard("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEnableWorkerCardSucces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryUpdateWorkerCardStatus).
		WithArgs(enableCardCode, "fname mname lname").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	lib := &Database{DB: db}

	if err = lib.EnableWorkerCard("fname mname lname"); err != nil {
		t.Errorf("error when disable card: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when worker not found
func TestEnableWorkerCardRollbackOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.EnableWorkerCard("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// failed when card disable
func TestEnableWorkerCardRollbackTwo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	column := []string{"ID"}
	rows := sqlmock.NewRows(column).AddRow("1").AddRow("2")

	mock.ExpectBegin()
	mock.ExpectQuery(testQueryFindWorkerIDByName).
		WithArgs("Fname", "Mname", "Lname").
		WillReturnRows(rows)
	mock.ExpectExec(testQueryUpdateWorkerCardStatus).
		WithArgs(enableCardCode, "fname mname lname").
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	lib := &Database{DB: db}

	if err = lib.EnableWorkerCard("fname mname lname"); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
