package libgorion

// Company model
type Company struct {
	Name         string
	WorkersCount uint
}

// Company get all comanies from Database
// return pionter to Company struct and error
func (db *Database) Company() ([]*Company, error) {
	rows, err := db.Query(queryCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies = make([]*Company, 0)
	for rows.Next() {
		c := new(Company)
		if err = rows.Scan(&c.Name, &c.WorkersCount); err != nil {
			return nil, err
		}

		companies = append(companies, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
