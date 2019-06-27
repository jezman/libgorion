package libgorion

import (
	"database/sql"
	"fmt"
	"strings"
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
func (db *Database) Events(firstDate, lastDate, worker string, door uint, denied bool) ([]*Event, error) {
	worker = strings.Title(worker)

	// change the query depending on the input flag
	switch {
	case door != 0 && worker != "":
		if err := validationWorker(worker); err != nil {
			return nil, err
		}
		rows, err = db.Query(queryEventsByEmployeeAndDoor, firstDate, lastDate, worker, door)
	case worker != "":
		if err := validationWorker(worker); err != nil {
			return nil, err
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
		e := new(Event)
		err = rows.Scan(
			&e.Worker.LastName,
			&e.Worker.FirstName,
			&e.Worker.MidName,
			&e.Worker.Company.Name,
			&e.FirstTime,
			&e.Action,
			&e.Door.Name,
		)
		if err != nil {
			return nil, err
		}

		e.Worker.FullName = joinNames(
			e.Worker.LastName,
			e.Worker.FirstName,
			e.Worker.MidName,
		)

		e.WorkedTime = e.LastTime.Sub(e.FirstTime)

		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// WorkedTime gets the list of workers and
// calculates their worked time
// return pointer to Event struct and error
func (db *Database) WorkedTime(firstDate, lastDate, worker, company string) ([]*Event, error) {
	if err = validationDate(firstDate); err != nil {
		return nil, err
	}

	if err = validationDate(lastDate); err != nil {
		return nil, err
	}

	worker = strings.Title(worker)

	switch {
	case worker != "":
		if err = validationWorker(worker); err != nil {
			return nil, err
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
		e := new(Event)
		err = rows.Scan(
			&e.Worker.LastName,
			&e.Worker.FirstName,
			&e.Worker.MidName,
			&e.Worker.Company.Name,
			&e.FirstTime,
			&e.LastTime,
		)

		if err != nil {
			return nil, err
		}

		e.Worker.FullName = joinNames(
			e.Worker.LastName,
			e.Worker.FirstName,
			e.Worker.MidName,
		)

		e.WorkedTime = e.LastTime.Sub(e.FirstTime)

		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// EventsValues return pointer to Event struct and error
func (db *Database) EventsValues() ([]*Event, error) {
	rows, err := db.Query(queryEventsValues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values = make([]*Event, 0)
	for rows.Next() {
		v := new(Event)
		if err = rows.Scan(&v.ID, &v.Action, &v.Description); err != nil {
			return nil, err
		}

		values = append(values, v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

// EventsTail puts tail events to STDOUT.
func (db *Database) EventsTail(interval time.Duration, worker string) error {
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
		e := new(Event)
		err := rows.Scan(
			&e.Worker.LastName,
			&e.Worker.FirstName,
			&e.Worker.MidName,
			&e.Worker.Company.Name,
			&e.FirstTime,
			&e.Action,
			&e.Door.Name,
		)

		if err != nil {
			return err
		}

		e.Worker.FullName = joinNames(
			e.Worker.LastName,
			e.Worker.FirstName,
			e.Worker.MidName,
		)

		fmt.Println(
			e.FirstTime.Format("02.01.2006 15:04:05"),
			e.Door.Name,
			ColorizedDenied(e.Action),
			e.Worker.Company.Name,
			colorizedWorker(e.Worker.FullName, worker),
		)
	}
	defer rows.Close()

	return nil
}
