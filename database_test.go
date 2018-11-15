package libgorion

import (
	"os"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func TestOpenDB(t *testing.T) {
	db, err := OpenDB(os.Getenv("BOLID_DSN"))
	if err != nil {
		t.Fatalf("error '%s' was not expected when opening a Database connection", err)
	}
	defer db.Close()
}
