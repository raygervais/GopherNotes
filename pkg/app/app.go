package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/raygervais/gophernotes/pkg/cli"
	"github.com/raygervais/gophernotes/pkg/conf"
	"github.com/raygervais/gophernotes/pkg/db"
)

var (
	helpFlag = flag.Bool("help", false, "Display application usage material")
)

// Application is a wrapper around main function so we can test various args
func Application() (int, string) {
	// Determine where to store the database for the user based on operating system
	configPath, err := conf.DetermineStorageLocation()
	if err != nil {
		return 1, fmt.Sprintf("Could not determine configuration location: %s", err)
	}

	if err := conf.InitializeConfigurationLocation(configPath); err != nil {
		return 1, fmt.Sprintf("Could not initialize configuration location: %s", err)
	}

	// Create and connect to database
	db, err := db.CreateDatabaseConnection(configPath +
		conf.ApplicationName + conf.DatabaseLocation)

	if err != nil {
		return 1, fmt.Sprint(err)
	}

	if err := db.InitializeNotesTable(); err != nil {
		return 1, fmt.Sprintf("Could not initialize database and tables: %s", err)
	}

	cli := cli.InitCLI(db)

	flag.Parse()

	if *helpFlag || len(os.Args) == 1 {
		cli.Help()
		return 0, ""
	}

	if err := cli.Handler(); err != nil {
		return 1, fmt.Sprintf("Command error: %v\n", err)
	}

	return 0, ""
}
