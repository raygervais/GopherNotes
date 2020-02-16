package main

import (
	"os"
	"testing"
)

const (
	TestFile        = "./test.logs"
	NonExistentFile = "./100WaysPHPIsBetter.md"
	ExistentFile    = "./README.md"
)

// Helper Functions
func CreateTemporaryTestFile(t *testing.T) {
	file := CreateNotesFile(TestFile)
	defer file.Close()
}

func DeleteTemporaryTestFile(t *testing.T) {
	err := os.Remove(TestFile)
	if err != nil {
		t.Errorf("Error in Deletion of Temporary Test File %v", err)
	}
}

// Test Cases
func TestFileOperations(t *testing.T) {
	t.Run("path validation should return false for non-existent file", func(t *testing.T) {
		want := false
		got := ValidateFilePath(NonExistentFile)

		AssertExpectedResultsBool(t, got, want)
	})

	t.Run("path validation should return true for existent file", func(t *testing.T) {
		want := true
		got := ValidateFilePath(ExistentFile)

		AssertExpectedResultsBool(t, got, want)
	})

	t.Run("should be able to create notes.log file", func(t *testing.T) {
		CreateTemporaryTestFile(t)
		defer DeleteTemporaryTestFile(t)
		want := true
		got := ValidateFilePath(TestFile)

		AssertExpectedResultsBool(t, got, want)

	})
	//		data := "testing: \"This is a test script\""
	//		written := WriteFile(file, []byte(data))
	//      written.Close()
}
