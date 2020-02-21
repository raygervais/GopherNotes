package main

import (
	"flag"
	"fmt"
	"os"
)

func GenerateFlag(argument string) *flag.FlagSet {
	return flag.NewFlagSet(argument, flag.ExitOnError)
}

func GenerateFlagParams(flag *flag.FlagSet, flagType, value, hint string) *string {
	return flag.String(flagType, value, hint)
}

// Verify we are providing a subcomamnd
// os.Arg[0] is the main command
// os.Arg[1] is the subcommand
func VerifyFlagArguments(arguments []string) {
	if len(arguments) < 2 {
		fmt.Println("Please provide a command, new or search")
		os.Exit(1)
	}
}

func VerifyFlagParams(param *string, flag *flag.FlagSet) {
	if *param == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
