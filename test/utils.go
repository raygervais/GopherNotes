package test

import (
	"testing"
)

// ExpectToEqualInt compares the equality of two ints provided.
func ExpectToEqualInt(t *testing.T, res, exp int) {
	if res != exp {
		t.Errorf("Expected: %d, Received: %d", exp, res)
		t.FailNow()
	}
}

// ExpectToEqualString compares the equality of two strings provided.
func ExpectToEqualString(t *testing.T, res, exp string) {
	if res != exp {
		t.Errorf("Expected: %s, Received: %s", exp, res)
		t.FailNow()
	}
}
