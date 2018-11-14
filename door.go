package libgorion

// Door model
type Door struct {
	ID   int
	Name string
}

// Doors get all doors and IDs from database
// return pionter to Door struct and error
func (db *database) Doors() ([]*Door, error) {
	rows, err := db.Query(queryDoors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var doors = make([]*Door, 0)
	for rows.Next() {
		d := new(Door)
		if err = rows.Scan(&d.ID, &d.Name); err != nil {
			return nil, err
		}

		doors = append(doors, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return doors, nil
}
