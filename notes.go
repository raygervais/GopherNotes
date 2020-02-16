package main

import (
	"strings"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

type Note struct {
	entry string
	date  string
}

// Context: Formatting
func CreateTimeStampFormat() string {
	return time.Now().Format(layoutISO)
}

// fx: a, b -> c
func CreateNoteEntry(timestamp, note string) Note {
	return Note{
		date:  timestamp,
		entry: note,
	}
}

// fx: a -> [a]
func ParseNoteEntry(note string) []string {
	return strings.Split(note, ":")
}

// Context: Filtering

// fx: [a], b -> [a]
func CreateSearchResults(notes []Note, text string) []Note {
	return Filter(notes, func(note Note) bool {
		return strings.Contains(note.entry, text)
	})
}

// Context: CRUD
// fx: a, a -> a
func UpdateNote(archived, entered string) Note {
	parsedNote := ParseNoteEntry(archived)
	parsedDate := parsedNote[0]

	return CreateNoteEntry(parsedDate, entered)
}
