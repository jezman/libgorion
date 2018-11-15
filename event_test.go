package libgorion

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var (
	timeNow     = time.Now().Local()
	firstDate   = timeNow.Format("02.01.2006")
	lastDate    = timeNow.AddDate(0, 0, 1).Format("02.01.2006")
	workerName  = "Worker"
	companyName = "Company"
	doorID      = uint(22)
	denied      = true
)

func TestEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Door", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", time.Now(), "door", "action")

	mock.ExpectQuery(testQueryEvents).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err := lib.Events(firstDate, lastDate, "", 0, false); err != nil {
		t.Errorf("error was not expected while gets all events %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByEmployeeAndDoor).
		WithArgs(firstDate, lastDate, workerName, doorID).
		WillReturnRows(rows)

	if _, err = lib.Events(firstDate, lastDate, workerName, doorID, false); err != nil {
		t.Errorf("error was not expected while gets events by worker and door %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByEmployee).
		WithArgs(firstDate, lastDate, workerName).
		WillReturnRows(rows)

	if _, err = lib.Events(firstDate, lastDate, workerName, 0, false); err != nil {
		t.Errorf("error was not expected while gets events by worker %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByDoor).
		WithArgs(firstDate, lastDate, doorID).
		WillReturnRows(rows)

	if _, err = lib.Events(firstDate, lastDate, "", doorID, false); err != nil {
		t.Errorf("error was not expected while gets events by door %q ", err)
	}

	mock.ExpectQuery(testQueryEventsDenied).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err = lib.Events(firstDate, lastDate, "", 0, denied); err != nil {
		t.Errorf("error was not expected while gets denied events %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestWorkedTime(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", timeNow, timeNow)

	mock.ExpectQuery(testQueryWorkedTime).WillReturnRows(rows)

	if _, err = lib.WorkedTime(firstDate, lastDate, "", ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(testQueryWorkedTimeByCompany).WillReturnRows(rows)

	if _, err = lib.WorkedTime(firstDate, lastDate, "", companyName); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(testQueryWorkedTimeByEmployee).WillReturnRows(rows)

	if _, err = lib.WorkedTime(firstDate, lastDate, workerName, ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEventsValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"ID", "Value", "Comment"}
	rows := sqlmock.NewRows(column).
		AddRow("1", "alert", "alert event")

	mock.ExpectQuery(testQueryEventsValues).WillReturnRows(rows)

	if _, err := lib.EventsValues(); err != nil {
		t.Errorf("error was not expected while gets events values %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEventsTail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub Database connection", err)
	}
	defer db.Close()

	lib := &Database{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Door", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", time.Now(), "door", "action")

	interval := time.Duration(5)
	timeNow := time.Now().Local()
	backForSeconds := timeNow.Add(time.Second * -interval)

	mock.ExpectQuery(testQueryEvents).
		WithArgs(backForSeconds.Format("02.01.2006 15:04:05"), timeNow.Format("02.01.2006 15:04:05")).
		WillReturnRows(rows)

	if err := lib.EventsTail(interval, ""); err != nil {
		t.Errorf("error was not expected while gets all events %q ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
