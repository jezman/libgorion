package libgorion

import (
	"testing"

	"github.com/bclicn/color"
)

func TestColorizedDenied(t *testing.T) {
	var testCases = []struct {
		input string
		want  string
	}{
		{"string", "string"},
		{"any_string", "any_string"},
		{"начало отклонен конец", color.Red("начало отклонен конец")},
		{"начало Запрет конец", color.Red("начало Запрет конец")},
	}
	for _, test := range testCases {
		if got := ColorizedDenied(test.input); got != test.want {
			t.Errorf("Event(%q) is %q. Need %q", test.input, got, test.want)
		}
	}
}

func TestColorizedWorker(t *testing.T) {
	var testCases = []struct {
		fullName string
		substr   string
		want     string
	}{
		{"Иванов", "Иванов", color.Yellow("Иванов")},
		{"Иванов", "иванов", color.Yellow("Иванов")},
		{"Петров", "Иванов", "Петров"},
		{"Петров", "", "Петров"},
	}
	for _, test := range testCases {
		if got := colorizedWorker(test.fullName, test.substr); got != test.want {
			t.Errorf("Worker(%q) is %q. Need %q", test.fullName, got, test.want)
		}
	}

}

func TestSplitFullName(t *testing.T) {
	var testCases = []struct {
		fullName string
		want     []string
	}{
		{"Иванов Иван Иванович", []string{"Иванов", "Иван", "Иванович"}},
		{" иванов иван иванович", []string{"Иванов", "Иван", "Иванович"}},
		{"Иванов Иван ", nil},
		{" Иван ", nil},
	}

	for _, test := range testCases {
		if got, _ := splitFullName(test.fullName); len(got) != len(test.want) {
			t.Errorf("length error: got - %v, want - %v", got, test.want)
		}
	}
}

func TestJoinNames(t *testing.T) {
	var testCases = []struct {
		firstName string
		midName   string
		lastName  string
		want      string
	}{
		{"Иванов", "Иван", "Иванович", "Иванов Иван Иванович"},
	}

	for _, c := range testCases {
		if got := joinNames(c.firstName, c.midName, c.lastName); len(got) != len(c.want) {
			t.Errorf("length error: got - %v, want - %v", got, c.want)
		}
	}
}
