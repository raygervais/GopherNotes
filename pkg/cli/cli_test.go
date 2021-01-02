package cli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/raygervais/gophernotes/pkg/cli"
	"github.com/raygervais/gophernotes/pkg/db"
	"github.com/raygervais/gophernotes/test"
)

var testCLI cli.CommandLineInterface

func setup() {
	fmt.Println("Setting up Database")
	db := db.CreateDatabaseConnection("/tmp/test.db")
	db.InitializeNotesTable()
	testCLI = cli.InitCLI(db)
}

func teardown() {
	fmt.Println("Tearing down: Deleting Database")
	os.Remove("/tmp/test.db")
}

func TestCrudInterface(t *testing.T) {
	setup()
	defer teardown()
	// First entry in command string array is empty since this would be the binary name during runtime
	// Ex. gn create --help
	//    [0] [1]   [2]
	testCases := []struct {
		desc    string
		command []string
		args    []string
		error   string
		code    int
		stdin   string
	}{
		// Create
		{
			desc:    "Create Valid Note",
			command: []string{"", "create"},
			args: []string{
				"--note",
				"'Testing 01'",
			},
			code: 0,
		},
		{
			desc:    "Create Invalid Note (Missing Args)",
			command: []string{"", "create"},
			args:    []string{},
			error:   "create expects at least 1 args, 0 provided.",
			code:    1,
		},

		// Fetch
		{
			desc:    "Fetch Valid",
			command: []string{"", "fetch"},
			args:    []string{},
			error:   "",
			code:    0,
		},
		{
			desc:    "Fetch Valid (Redundant Args)",
			command: []string{"", "fetch"},
			args:    []string{"Maybe fetch a few?"},
			error:   "",
			code:    0,
		},

		// Search
		{
			desc:    "Search Valid",
			command: []string{"", "search"},
			args:    []string{"--note", "Testing"},
			error:   "",
			code:    0,
		},
		{
			desc:    "Search Invalid (Missing Args)",
			command: []string{"", "search"},
			args:    []string{},
			error:   "search expects at least 1 args, 0 provided.",
			code:    1,
		},

		// Delete
		{
			desc:    "Delete Valid",
			command: []string{"", "delete"},
			args:    []string{"--id", "1", "--y"},
			error:   "",
			code:    0,
		},

		{
			desc:    "Delete Invalid",
			command: []string{"", "delete"},
			args:    []string{"--id", "-101023", "--y"},
			error:   "No rows were deleted with rowid: -101023\n",
			code:    1,
		},

		// Help
		{
			desc:    "Help Valid",
			command: []string{"", "--help"},
			args:    []string{},
			error:   "",
			code:    0,
		},

		{
			desc:    "Help Valid With Arguments",
			command: []string{"", "--help"},
			args:    []string{"--create"},
			error:   "",
			code:    0,
		},

		// CLI Error Handling
		{
			desc:    "Invalid Command",
			command: []string{"", "testing"},
			args:    []string{},
			error:   "Invalid command provided: 'testing'",
			code:    1,
		},

		{
			desc:    "Misspelled Command",
			command: []string{"", "creat"},
			args:    []string{},
			error:   "Invalid command provided: 'creat'",
			code:    1,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Helper()

			os.Args = append(tC.command, tC.args...)
			err := testCLI.Handler()

			if err != nil && tC.error != "" {
				t.Log(err)
				test.ExpectToEqualInt(t, tC.code, 1)
				test.ExpectToEqualString(t, tC.error, err.Error())
			}

		})
	}
}
