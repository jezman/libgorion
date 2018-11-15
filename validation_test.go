package libgorion

import (
	"errors"
	"testing"
)

func TestValidationWorkers(t *testing.T) {
	validCases := []string{"user", "сотрудник"}
	invalidCases := []string{"a", "123456"}

	for _, test := range validCases {
		if err := validationWorker(test); err != nil {
			t.Errorf("regexp don't match: (%v)", err)
		}
	}

	for _, test := range invalidCases {
		err := validationWorker(test)
		if err == errors.New("invalid worker. allowed only letters") {
			t.Errorf("regexp don't match: (%v)", err)
		}
	}
}

func TestValidationDate(t *testing.T) {
	validCases := []string{"02.12.2007", "02-12-2007"}
	invalidCases := []string{"02/12/2007", "f", "123456"}

	for _, test := range validCases {
		if err := validationDate(test); err != nil {
			t.Errorf("regexp don't match: (%v)", err)
		}

		for _, test := range invalidCases {
			err = validationDate(test)
			if err == errors.New("invalid date. corrects format: DD.MM.YYYY or DD-MM-YYYY") {
				t.Errorf("regexp don't match: (%v)", err)
			}
		}
	}

}
