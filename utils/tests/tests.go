package tests

import (
	"testing"
)

// AssertEquals verify that a test is in error
func AssertEquals(t *testing.T, context string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s test fail; expected %d; actual %d", context, expected, actual)
	}
}
