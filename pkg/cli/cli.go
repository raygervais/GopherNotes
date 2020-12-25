package cli

import (
	"fmt"
	"os"
)

type CommandLineInterface struct {
	commands map[string]func() func(string) error
}

func InitCLI() CommandLineInterface {
	cli := CommandLineInterface{}

	cli.commands = map[string]func() func(string) error{
		"create": cli.create,
		//		"edit":   cli.edit,
		//		"fetch":  cli.fetch,
		//		"delete": cli.delete,
		//		"search": cli.search,
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
		return nil
	}
}
