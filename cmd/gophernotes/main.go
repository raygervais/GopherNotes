package main

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

func main() {
	// Determine where to store the database for the user based on operating system
	configPath, err := conf.DetermineStorageLocation()
	if err != nil {
		fmt.Printf("Could not determine configuration location: %s", err)
		os.Exit(1)
	}

	if err := conf.InitializeConfigurationLocation(configPath); err != nil {
		fmt.Printf("Could not initialize configuration location: %s", err)
		os.Exit(1)
	}

	// Create and connect to database
	db := db.CreateDatabaseConnection(configPath +
		conf.ApplicationName + conf.DatabaseLocation)
	if err := db.InitializeNotesTable(); err != nil {
		fmt.Printf("Could not initialize database and tables: %s", err)
		os.Exit(1)
	}

	cli := cli.InitCLI(db)

	flag.Parse()

	if *helpFlag || len(os.Args) == 1 {
		cli.Help()
		return
	}

	if err := cli.Handler(); err != nil {
		fmt.Printf("Command error: %v\n", err)
		os.Exit(1)
	}
}
