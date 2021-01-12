package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/raygervais/gophernotes/pkg/db"
)

// TestPath implies the location of our TempDir to store configuration and DB
var TestPath = os.TempDir() + "/test.db"

// ExpectToEqualInt compares the equality of two ints provided.
func ExpectToEqualInt(t *testing.T, res, exp int) {
	if res != exp {
		t.Errorf("Expected: %d, Received: %d", exp, res)
		t.FailNow()
	}
}

// ExpectToEqualString compares the equality of two strings provided.
func ExpectToEqualString(t *testing.T, res, exp string) {
	if res != exp {
		t.Errorf("Expected: %s, Received: %s", exp, res)
		t.FailNow()
	}
}

// ExpectToContain verifies that a substring exists within the provided string.
func ExpectToContain(t *testing.T, res, exp string) {
	if !strings.Contains(res, exp) {
		t.Errorf("Expected %s to contain %s", res, exp)
		t.FailNow()
	}
}

// SetupDatabase orchestrates the creation of our testDB found in os.TempDir
func SetupDatabase(t *testing.T) db.Database {
	fmt.Println("Setting up Database")
	db, err := db.CreateDatabaseConnection(TestPath)
	if err != nil {
		t.Error(err)
	}

	db.InitializeNotesTable()
	return db
}

// TeardownDatabase orchestrates deletion of our testDB found in os.TempDir
func TeardownDatabase(t *testing.T) {
	fmt.Println("Tearing down: Deleting Database")
	if err := os.Remove(TestPath); err != nil {
		t.Error(err)
	}
}
