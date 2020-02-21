package main

import "database/sql"

type Note struct {
	entry string
	date  string
	id    int
}

func ParseDatabaseRows(rows *sql.Rows) []Note {
	notes := []Note{}
	for rows.Next() {
		note := Note{}
		rows.Scan(&note.id, &note.entry, &note.date)
		notes = append(notes, note)
	}

	return notes
}

func GenerateNoteOutput(note Note) string {
	return note.date + ": " + note.entry
}

func PrintNoteOutput(notes []Note) []string {
	output := []string{}

	for _, note := range notes {
		output = append(output, GenerateNoteOutput(note))
	}

	return output
}
