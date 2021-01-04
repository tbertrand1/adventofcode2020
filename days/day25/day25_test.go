package main

import (
	"../../utils/tests"
	"testing"
)

func TestPart1WithTestInputs(t *testing.T) {
	tests.AssertEquals(t, 14897079, part1(5764801, 17807724))
}

func TestPart1(t *testing.T) {
	tests.AssertEquals(t, 4441893, part1(inputCardPublicKey, inputDoorPublicKey))
}
