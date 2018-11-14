package libgorion

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var (
	rows  *sql.Rows
	err   error
	query string
)

// Event model
type Event struct {
	Worker      Worker
	FirstTime   time.Time
	LastTime    time.Time
	Company     Company
	Door        Door
	Action      string
	Description string
	ID          string
	WorkedTime  time.Duration
}

// Events gets the list of events for the time period
// return pointer to Event struct and error
func (db *DB) Events(firstDate, lastDate, worker string, door uint, denied bool) ([]*Event, error) {
	// change the query depending on the input flag
	switch {
	case door != 0 && worker != "":
		if !validationEmployee(worker) {
			fmt.Print("invalid worker. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(queryEventsByEmployeeAndDoor, firstDate, lastDate, worker, door)
	case worker != "":
		if !validationEmployee(worker) {
			fmt.Print("invalid worker. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(queryEventsByEmployee, firstDate, lastDate, worker)
	case door != 0:
		rows, err = db.Query(queryEventsByDoor, firstDate, lastDate, door)
	case denied:
		rows, err = db.Query(queryEventsDenied, firstDate, lastDate)
	default:
		rows, err = db.Query(queryEvents, firstDate, lastDate)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err = rows.Scan(
			&event.Worker.LastName,
			&event.Worker.FirstName,
			&event.Worker.MidName,
			&event.Worker.Company.Name,
			&event.FirstTime,
			&event.Action,
			&event.Door.Name,
		)
		if err != nil {
			return nil, err
		}

		event.Worker.FullName = event.Worker.LastName + " " +
			event.Worker.FirstName + " " + event.Worker.MidName

		event.WorkedTime = event.LastTime.Sub(event.FirstTime)

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil

}

// WorkedTime gets the list of workers and
// calculates their worked time
// return pointer to Event struct and error
func (db *DB) WorkedTime(firstDate, lastDate, worker, company string) ([]*Event, error) {
	if !validationDate(firstDate) || !validationDate(lastDate) {
		fmt.Print("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY")
		os.Exit(1)
	}

	switch {
	case worker != "":
		if !validationEmployee(worker) {
			fmt.Print("invalid worker. allowed only letters")
			os.Exit(1)
		}
		rows, err = db.Query(queryWorkedTimeByEmployee, firstDate, lastDate, worker)
	case company != "":
		rows, err = db.Query(queryWorkedTimeByCompany, firstDate, lastDate, company)
	default:
		rows, err = db.Query(queryWorkedTime, firstDate, lastDate)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]*Event, 0)
	for rows.Next() {
		event := new(Event)
		err = rows.Scan(
			&event.Worker.LastName,
			&event.Worker.FirstName,
			&event.Worker.MidName,
			&event.Worker.Company.Name,
			&event.FirstTime,
			&event.LastTime,
		)

		if err != nil {
			return nil, err
		}

		event.Worker.FullName = event.Worker.LastName + " " +
			event.Worker.FirstName + " " + event.Worker.MidName

		event.WorkedTime = event.LastTime.Sub(event.FirstTime)

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// EventsValues return pointer to Event struct and error
func (db *DB) EventsValues() ([]*Event, error) {
	rows, err := db.Query(queryEventsValues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventValues = make([]*Event, 0)
	for rows.Next() {
		ev := new(Event)
		if err = rows.Scan(&ev.ID, &ev.Action, &ev.Description); err != nil {
			return nil, err
		}

		eventValues = append(eventValues, ev)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return eventValues, nil
}

// EventsTail puts tail events to STDOUT.
func (db *DB) EventsTail(interval time.Duration, worker string) error {
	timeNow := time.Now().Local()
	backForSeconds := timeNow.Add(time.Second * -interval)

	rows, err := db.Query(
		queryEvents,
		backForSeconds.Format("02.01.2006 15:04:05"),
		timeNow.Format("02.01.2006 15:04:05"),
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		event := new(Event)
		err := rows.Scan(
			&event.Worker.LastName,
			&event.Worker.FirstName,
			&event.Worker.MidName,
			&event.Worker.Company.Name,
			&event.FirstTime,
			&event.Action,
			&event.Door.Name,
		)

		if err != nil {
			return err
		}

		event.Worker.FullName = event.Worker.LastName + " " +
			event.Worker.FirstName + " " + event.Worker.MidName

		fmt.Println(
			event.FirstTime.Format("02.01.2006 15:04:05"),
			event.Door.Name,
			colorizedDenied(event.Action),
			event.Worker.Company.Name,
			colorizedWorker(event.Worker.FullName, worker),
		)
	}
	defer rows.Close()

	return nil
}
