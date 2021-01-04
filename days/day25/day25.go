package main

import (
	"fmt"
)

const inputCardPublicKey = 16616892
const inputDoorPublicKey = 14505727

func main() {
	fmt.Printf("Part 1: %d\n", part1(inputCardPublicKey, inputDoorPublicKey))
}

func part1(cardPublicKey int, doorPublicKey int) int {
	cardLoopSize := 0

	value := 1
	for value != cardPublicKey {
		value = (value * 7) % 20201227
		cardLoopSize += 1
	}

	encryptionKey := 1
	for n := 0; n < cardLoopSize; n++ {
		encryptionKey = (encryptionKey * doorPublicKey) % 20201227
	}

	return encryptionKey
}
