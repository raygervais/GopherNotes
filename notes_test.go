package main

import (
	"testing"
	"time"
)

var timeStamp = time.Now().Format(layoutISO)

func TestNotesHandling(t *testing.T) {
	t.Run("Should create appropriate format", func(t *testing.T) {
		want := Note{date: timeStamp, entry: "Testing Note"}
		got := CreateNoteEntry(timeStamp, "Testing Note")

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Timestamp shoud be in ISO format", func(t *testing.T) {
		want := timeStamp
		got := CreateTimeStampFormat()

		AssertExpectedResultsString(t, got, want)
	})

	t.Run("Should split entry into two array items, a date and text", func(t *testing.T) {
		want := []string{
			"2020-02-14",
			"\"Hello, World!\"",
		}

		got := ParseNoteEntry("2020-02-14:\"Hello, World!\"")
		AssertExpectedResultsStringArray(t, got, want)
	})

	t.Run("Should update entry with new text and old date", func(t *testing.T) {
		oldNote := "2020-02-14:\"Hello, World!\""
		want := Note{date: "2020-02-14", entry: "Hello, Open Source!"}
		got := UpdateNote(oldNote, "Hello, Open Source!")

		AssertExpectedResultsNotes(t, got, want)
	})
}
