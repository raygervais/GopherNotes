package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/raygervais/gophernotes/pkg/db"
	"github.com/raygervais/gophernotes/test"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestDatabaseConnection(t *testing.T) {
	if _, err := os.Stat(test.TestPath); os.IsExist(err) {
		test.TeardownDatabase(t)
	}

	db, err := db.CreateDatabaseConnection(test.TestPath)
	if err != nil {
		t.Error(err)
	}

	if err := db.InitializeNotesTable(); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(test.TestPath); os.IsNotExist(err) {
		t.Error(err)
	}

	test.TeardownDatabase(t)
}

func TestDatabaseCreate(t *testing.T) {
	db := test.SetupDatabase(t)
	defer test.TeardownDatabase(t)

	t.Run("Create new note", func(t *testing.T) {
		err := db.Create("Testing Database")
		if err != nil {
			t.Fail()
		}
	})

	t.Run("Create invalid empty note", func(t *testing.T) {
		err := db.Create("")
		test.ExpectToEqualString(t, err.Error(), "Invalid input provided as message parameter")
	})
}

func TestDatabaseFetch(t *testing.T) {
	db := test.SetupDatabase(t)

	for i := 0; i < 10; i++ {
		db.Create(fmt.Sprintf("TestFetchFunction%d", i))
	}

	defer test.TeardownDatabase(t)

	t.Run("Fetch all database entries", func(t *testing.T) {
		_, err := db.Fetch(10, "asc")
		if err != nil {
			t.Fail()
		}
	})

	t.Run("Fetch only 1 entry", func(t *testing.T) {
		res, err := db.Fetch(1, "asc")
		if err != nil {
			t.Fail()
		}

		cmp := dbRowIterateMock(t, res)

		test.ExpectToEqualInt(t, len(cmp), 1)
		test.ExpectToContain(t, cmp[0], "TestFetchFunction0")
	})

	t.Run("Fetch only 1 with sort of desc", func(t *testing.T) {
		res, err := db.Fetch(1, "desc")
		if err != nil {
			t.Fail()
		}
		cmp := dbRowIterateMock(t, res)

		test.ExpectToEqualInt(t, len(cmp), 1)
		test.ExpectToContain(t, cmp[0], "TestFetchFunction9")
	})
}

func TestDatabaseEdit(t *testing.T) {
	db := test.SetupDatabase(t)

	for i := 0; i < 10; i++ {
		db.Create(fmt.Sprintf("TestFetchFunction%d", i))
	}

	defer test.TeardownDatabase(t)

	t.Run("Edit valid entry", func(t *testing.T) {
		err := db.EditByID(1, "TestEditFunction1")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Edit invalid entry with missing edit", func(t *testing.T) {
		err := db.EditByID(100, "")
		test.ExpectToEqualString(t, "Invalid input provided as change parameter", err.Error())
	})

	t.Run("Edit invalid entry with missing id", func(t *testing.T) {
		err := db.EditByID(-1, "")
		test.ExpectToEqualString(t, "Invalid input provided as id parameter", err.Error())
	})
}

func TestDatabaseSearch(t *testing.T) {
	db := test.SetupDatabase(t)

	for i := 0; i < 10; i++ {
		db.Create(fmt.Sprintf("TestFetchFunction%d", i))
	}

	defer test.TeardownDatabase(t)

	t.Run("Search valid entry", func(t *testing.T) {
		res, err := db.Search("TestFetchFunction1", 1, "asc")
		if err != nil {
			t.Error(err)
		}

		cmp := dbRowIterateMock(t, res)

		test.ExpectToEqualInt(t, len(cmp), 1)
		test.ExpectToContain(t, cmp[0], "TestFetchFunction1")
	})

	t.Run("Search invalid entry with missing message", func(t *testing.T) {
		_, err := db.Search("", 1, "asc")
		test.ExpectToEqualString(t, "Invalid input provided as message parameter", err.Error())
	})
}

func TestDatabaseDelete(t *testing.T) {
	db := test.SetupDatabase(t)

	for i := 0; i < 10; i++ {
		db.Create(fmt.Sprintf("TestFetchFunction%d", i))
	}

	defer test.TeardownDatabase(t)

	t.Run("Delete valid entry", func(t *testing.T) {
		err := db.DeleteByID(1)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Delete invalid entry with missing id", func(t *testing.T) {
		err := db.DeleteByID(-1)
		test.ExpectToEqualString(t, "Invalid input provided as id parameter", err.Error())
	})

	t.Run("Delete invalid entry with non-existent id", func(t *testing.T) {
		err := db.DeleteByID(10000)
		test.ExpectToEqualString(t, "0 rows were deleted with rowid: 10000", err.Error())
	})
}

// Helper function, appends vs concats so we can count returned rows
func dbRowIterateMock(t *testing.T, res *sql.Rows) []string {
	var cmp []string

	var note, date, id string

	for res.Next() {
		if err := res.Scan(&id, &note, &date); err != nil {
			t.Error(err)
		}
		cmp = append(cmp, fmt.Sprintf("%v) %v: %v\n", id, date, note))
	}

	res.Close()

	return cmp
}
