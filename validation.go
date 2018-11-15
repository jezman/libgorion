package libgorion

import (
	"errors"
	"regexp"
)

const (
	workerFormat = `^[а-яА-Яa-zA-z][а-яa-z-А-ЯA-Z-_\.]{0,20}$`
	dateFormat   = `(0[1-9]|[12][0-9]|3[01])[- ..](0[1-9]|1[012])[- ..][201]\d\d\d`
)

// ValidationEmployee validation worker flag
// return error if value don't match regexp
func validationWorker(value string) error {
	match, _ := regexp.MatchString(workerFormat, value)
	if !match {
		return errors.New("invalid worker. allowed only letters")
	}

	return nil
}

// ValidationDate validation date flags
// return error if date don't match
// regexp DD.MM.YYYY or DD-MM-YYYY
func validationDate(date string) error {
	match, _ := regexp.MatchString(dateFormat, date)
	if !match {
		return errors.New("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY")
	}

	return nil
}
