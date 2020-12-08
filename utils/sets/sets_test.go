package sets

import (
	"../tests"
	"testing"
)

func TestSetOfStrings(t *testing.T) {
	set := NewSet()

	tests.AssertEquals(t, 0, set.Length())

	set.Add("Red")
	set.Add("Blue")

	tests.AssertEquals(t, 2, set.Length())
	tests.AssertTrue(t, set.Contains("Red"))
	tests.AssertFalse(t, set.Contains("Yellow"))

	set.Remove("Red")

	tests.AssertEquals(t, 1, set.Length())
	tests.AssertFalse(t, set.Contains("Red"))
}

func TestSetOfInts(t *testing.T) {
	set := NewSet()

	tests.AssertEquals(t, 0, set.Length())

	set.Add(1)
	set.Add(2)
	set.Add(2)

	tests.AssertEquals(t, 2, set.Length())
	tests.AssertTrue(t, set.Contains(1))
	tests.AssertTrue(t, set.Contains(2))
	tests.AssertFalse(t, set.Contains(3))

	set.Remove(2)

	tests.AssertEquals(t, 1, set.Length())
	tests.AssertFalse(t, set.Contains(2))
}
