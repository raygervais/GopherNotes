package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

var NOTES = []string{"2020-02-14:\"Hello, World!\"", "Hello Ray!"}

func CreateTimeStampFormat() string {
	return time.Now().Format(layoutISO)
}

// fx a,a -> b
func CreateNoteEntryText(timestamp, note string) string {
	return timestamp + ":\"" + note + "\""
}

func CreateSearchResults(notes []string, text string) []string {
	results := make([]string, 0)

	for _, note := range notes {
		if strings.Contains(note, text) {
			results = append(results, note)
		}
	}

	return results
}

// Application Entry Point
func main() {
	// Subcommand Parents
	newCommand := flag.NewFlagSet("new", flag.ExitOnError)
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)

	// New Note Command Flag Pointers
	newTextPtr := newCommand.String("text", "", "Note text to be saved")

	// Search Note Command Flag Pointers
	searchTextPtr := searchCommand.String("text", "", "Search notes with the following text")

	// Verify we are providing a subcomamnd
	// os.Arg[0] is the main command
	// os.Arg[1] is the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command, new or search")
		os.Exit(1)
	}

	// Switch on subcommand parsing
	switch os.Args[1] {
	case "new":
		newCommand.Parse(os.Args[2:])
	case "search":
		searchCommand.Parse(os.Args[2:])
	default:
		fmt.Println("Subargument Parsing failed!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Validate Required Fields
	if newCommand.Parsed() {
		// Required Flags
		if *newTextPtr == "" {
			newCommand.PrintDefaults()
			os.Exit(1)
		}

		// Call New Note Controller
		NOTES = append(NOTES,
			CreateNoteEntryText(
				CreateTimeStampFormat(), *newTextPtr))
		fmt.Println(strings.Join(NOTES, "\n"))
	}

	if searchCommand.Parsed() {
		// Required Flags
		if *searchTextPtr == "" {
			searchCommand.PrintDefaults()
			os.Exit(1)
		}

		// Call Search Note Controller
		note := CreateSearchResults(NOTES, *searchTextPtr)
		if len(note) == 0 {
			fmt.Println("Failed to find:", *searchTextPtr)
			os.Exit(1)
		}

		fmt.Println("Search Results:")
		fmt.Println(strings.Join(note, "\n"))
	}
}
