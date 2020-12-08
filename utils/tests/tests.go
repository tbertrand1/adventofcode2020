package tests

import (
	"testing"
)

func AssertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Test failed; expected to be true")
	}
}

func AssertFalse(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Test failed; expected to be false")
	}
}

func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Test failed; expected %d; actual %d", expected, actual)
	}
}
