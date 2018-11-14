package libgorion

import (
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

var (
	timeNow   = time.Now().Local()
	firstDate = timeNow.Format("02.01.2006")
	lastDate  = timeNow.AddDate(0, 0, 1).Format("02.01.2006")
	worker    = "Worker"
	company   = "Company"
	door      = uint(22)
	denied    = true
)

func TestEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Door", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", time.Now(), "door", "action")

	mock.ExpectQuery(testQueryEvents).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err := app.Events(firstDate, lastDate, "", 0, false); err != nil {
		t.Errorf("error was not expected while gets all events %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByEmployeeAndDoor).
		WithArgs(firstDate, lastDate, worker, door).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, worker, door, false); err != nil {
		t.Errorf("error was not expected while gets events by worker and door %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByEmployee).
		WithArgs(firstDate, lastDate, worker).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, worker, 0, false); err != nil {
		t.Errorf("error was not expected while gets events by worker %q ", err)
	}

	mock.ExpectQuery(testQueryEventsByDoor).
		WithArgs(firstDate, lastDate, door).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, "", door, false); err != nil {
		t.Errorf("error was not expected while gets events by door %q ", err)
	}

	mock.ExpectQuery(testQueryEventsDenied).
		WithArgs(firstDate, lastDate).
		WillReturnRows(rows)

	if _, err = app.Events(firstDate, lastDate, "", 0, denied); err != nil {
		t.Errorf("error was not expected while gets denied events %q ", err)
	}
}

func TestWorkedTime(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", timeNow, timeNow)

	mock.ExpectQuery(testQueryWorkedTime).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, "", ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(testQueryWorkedTimeByCompany).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, "", company); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}

	mock.ExpectQuery(testQueryWorkedTimeByEmployee).WillReturnRows(rows)

	if _, err = app.WorkedTime(firstDate, lastDate, worker, ""); err != nil {
		t.Errorf("error was not expected while gets worked time %q ", err)
	}
}

func TestEventsValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"ID", "Value", "Comment"}
	rows := sqlmock.NewRows(column).
		AddRow("1", "alert", "alert event")

	mock.ExpectQuery(testQueryEventsValues).WillReturnRows(rows)

	if _, err := app.EventsValues(); err != nil {
		t.Errorf("error was not expected while gets events values %q ", err)
	}
}

func TestEventsTail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a strub database connection", err)
	}
	defer db.Close()

	app := &DB{DB: db}

	column := []string{"Time", "firstName", "midName", "lastName", "Company", "Door", "Event"}
	rows := sqlmock.NewRows(column).
		AddRow("firstName", "midName", "lastName", "company", time.Now(), "door", "action")

	interval := time.Duration(5)
	timeNow := time.Now().Local()
	backForSeconds := timeNow.Add(time.Second * -interval)

	mock.ExpectQuery(testQueryEvents).
		WithArgs(backForSeconds.Format("02.01.2006 15:04:05"), timeNow.Format("02.01.2006 15:04:05")).
		WillReturnRows(rows)

	if err := app.EventsTail(interval, ""); err != nil {
		t.Errorf("error was not expected while gets all events %q ", err)
	}
}
