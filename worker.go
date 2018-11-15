package libgorion

import (
	"database/sql"
	"errors"
)

const (
	blockCardCode   = 32896
	unblockCardCode = 128
)

// Worker model
type Worker struct {
	FirstName string
	LastName  string
	MidName   string
	FullName  string
	Company   Company
}

// Workers get all workers from Database
// return pionter to Worker struct and error
func (db *Database) Workers(companyName string) ([]*Worker, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if companyName != "" {
		rows, err = db.Query(queryEmployeesByCompany, companyName)
	} else {
		rows, err = db.Query(queryEmployees)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workers = make([]*Worker, 0)
	for rows.Next() {
		w := new(Worker)
		if err = rows.Scan(&w.LastName, &w.FirstName, &w.MidName, &w.Company.Name); err != nil {
			return nil, err
		}

		w.FullName = joinNames(w.LastName, w.FirstName, w.MidName)

		workers = append(workers, w)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return workers, nil
}

// AddWorker to ACS
func (db *Database) AddWorker(name string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	fullName, err := splitFullName(name)
	if err != nil {
		return
	}

	if _, err = tx.Exec(queryAddWorker, fullName[0], fullName[1], fullName[2]); err != nil {
		return
	}

	return
}

// DeleteWorker from ACS Database
func (db *Database) DeleteWorker(name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if err = db.findWorkerIDByName(name); err != nil {
		return err
	}

	if _, err = tx.Exec(queryDeleteWorkerCards, name); err != nil {
		return err
	}

	fullName, err := splitFullName(name)
	if err != nil {
		return err
	}

	if _, err = tx.Exec(queryDeleteWorker, fullName[0], fullName[1], fullName[2]); err != nil {
		return err
	}

	return err
}

// DisableWorkerCard by worker full name.
func (db *Database) DisableWorkerCard(name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if err = db.findWorkerIDByName(name); err != nil {
		return err
	}

	if _, err = tx.Exec(queryUpdateWorkerCardStatus, blockCardCode, name); err != nil {
		return err
	}

	return err
}

// EnableWorkerCard by worker full name.
func (db *Database) EnableWorkerCard(name string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if err = db.findWorkerIDByName(name); err != nil {
		return err
	}

	if _, err = tx.Exec(queryUpdateWorkerCardStatus, unblockCardCode, name); err != nil {
		return err
	}

	return err
}

func (db *Database) findWorkerIDByName(names string) error {
	fullName, err := splitFullName(names)
	if err != nil {
		return err
	}

	rows, err = db.Query(queryFindWorkerIDByName, fullName[0], fullName[1], fullName[2])
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("worker not found")
	}

	return nil
}
