package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/raygervais/gophernotes/pkg/db"
	"github.com/raygervais/gophernotes/pkg/edit"
)

type CommandLineInterface struct {
	commands map[string]func() func(string) error
	database db.Database
}

func InitCLI(db db.Database) CommandLineInterface {
	cli := CommandLineInterface{
		database: db,
	}

	cli.commands = map[string]func() func(string) error{
		"create": cli.create,
		"edit":   cli.edit,
		"fetch":  cli.fetch,
		//		"delete": cli.delete,
		"search": cli.search,
		//		"health": cli.health,
	}

	return cli
}

func (cli CommandLineInterface) Handler() error {
	action := os.Args[1]

	// Lookup the command provided
	cmd, ok := cli.commands[action]
	if !ok {
		return fmt.Errorf("Invalid command provided: '%s'", action)
	}

	// Invoke the returned function
	return cmd()(action)
}

func (cli CommandLineInterface) Help() {
	var help string

	for name := range cli.commands {
		help += name + "\t --help\n"
	}

	fmt.Printf("Usage of %s:\n <command> [<args>]\n%s", os.Args[0], help)
}

func (cli CommandLineInterface) create() func(string) error {
	return func(cmd string) error {
		createCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		note := createCmd.String("note", "", "The note to store")

		if err := cli.checkArgs(1); err != nil {
			return err
		}

		if err := cli.parseCmd(createCmd); err != nil {
			return err
		}

		// Do DB Transaction
		if err := cli.database.Create(*note); err != nil {
			return err
		}

		return nil
	}
}

func (cli CommandLineInterface) edit() func(string) error {
	return func(cmd string) error {
		editCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		id := editCmd.Int("id", -1, "The note to store")

		if err := cli.checkArgs(1); err != nil {
			return err
		}

		if err := cli.parseCmd(editCmd); err != nil {
			return err
		}

		var note, date string
		res, err := cli.database.RetrieveByID(*id)
		if err != nil {
			return err
		}

		res.Scan(&note, &date)

		changes, err := edit.CaptureInputFromEditor(*id, note, date)
		if err != nil {
			return err
		}

		return cli.database.EditByID(*id, string(changes))
	}
}

func (cli CommandLineInterface) fetch() func(string) error {
	return func(cmd string) error {
		rows, err := cli.database.Fetch()
		if err != nil {
			return err
		}

		res, err := cli.database.IterateOnRows(rows)
		if err != nil {
			return err
		}

		fmt.Print(res)
		return nil
	}
}

func (cli CommandLineInterface) search() func(string) error {
	return func(cmd string) error {
		searchCmd := flag.NewFlagSet(cmd, flag.ExitOnError)
		note := searchCmd.String("note", "", "The note text to search")

		if err := cli.checkArgs(1); err != nil {
			return err
		}

		if err := cli.parseCmd(searchCmd); err != nil {
			return err
		}

		rows, err := cli.database.Search(*note)
		if err != nil {
			return err
		}

		res, err := cli.database.IterateOnRows(rows)
		if err != nil {
			return err
		}

		fmt.Print(res)
		return nil
	}
}

// Utility Functions

func (cli CommandLineInterface) parseCmd(cmd *flag.FlagSet) error {
	if err := cmd.Parse(os.Args[2:]); err != nil {
		return fmt.Errorf("Could not parse %s: %s", cmd.Name(), err)
	}

	return nil
}

func (cli CommandLineInterface) checkArgs(minArgs int) error {
	if len(os.Args) == 3 && os.Args[2] == "--help" {
		return nil
	}

	if len(os.Args)-2 < minArgs {
		fmt.Printf("Incorrect use of %s\n%s %s --help\n",
			os.Args[1], os.Args[0], os.Args[1])
		return fmt.Errorf("%s expects at least %d args, %d provided.",
			os.Args[1], minArgs, len(os.Args)-2)
	}

	return nil
}
