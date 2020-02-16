package main

import (
	"reflect"
	"strings"
	"testing"
)

var TestStrings = []string{"peach", "apple", "pear", "plum"}

// Helper Functions
func AssertExpectedResultsBool(t *testing.T, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertExpectedResultsString(t *testing.T, got, want []string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCollectionUtilities(t *testing.T) {
	t.Run("Index should return 2 for pear", func(t *testing.T) {
		want := 2
		got := Index(TestStrings, "pear")

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("Include should return false for non-existent", func(t *testing.T) {
		want := false
		got := Include(TestStrings, "grape")

		AssertExpectedResultsBool(t, got, want)
	})

	t.Run("Confirms any item in the list starts with 'P'", func(t *testing.T) {
		want := true
		got := Any(TestStrings, func(v string) bool {
			return strings.HasPrefix(v, "p")
		})

		AssertExpectedResultsBool(t, got, want)
	})

	t.Run("Confirms all item in the list does not start with 'P'", func(t *testing.T) {
		want := false
		got := All(TestStrings, func(v string) bool {
			return strings.HasPrefix(v, "p")
		})

		AssertExpectedResultsBool(t, got, want)
	})

	t.Run("Lists any item in the list contains 'e'", func(t *testing.T) {
		want := []string{"peach", "apple", "pear"}
		got := Filter(TestStrings, func(v string) bool {
			return strings.Contains(v, "e")
		})

		AssertExpectedResultsString(t, got, want)
	})

	t.Run("Returns array with transformed list items", func(t *testing.T) {
		want := []string{"PEACH", "APPLE", "PEAR", "PLUM"}
		got := Map(TestStrings, strings.ToUpper)

		AssertExpectedResultsString(t, got, want)
	})
}
