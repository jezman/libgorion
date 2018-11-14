package libgorion

import (
	"errors"
	"strings"

	"github.com/bclicn/color"
)

// ColorizedDenied events
func ColorizedDenied(event string) string {
	if strings.Contains(event, "отклонен") || strings.Contains(event, "Запрет") {
		return color.Red(event)
	}

	return event
}

// ColorizedWorker fullname
func colorizedWorker(str, substr string) string {
	if substr != "" && strings.Contains(str, strings.Title(substr)) {
		return color.Yellow(str)
	}

	return str
}

// SplitFullName return array of three elements: [first|mid|last]name
func splitFullName(name string) ([]string, error) {
	name = strings.Title(strings.Trim(name, " "))
	fullName := strings.Split(name, " ")

	if len(fullName) != 3 {
		return nil, errors.New("need first, mid and last name")
	}

	return fullName, nil
}

func joinNames(names ...string) string {
	return strings.Join(names, " ")
}
