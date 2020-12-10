package slices

// Find returns the index of a int in an array of ints
func Find(values []int, value int) (int, bool) {
	for i, val := range values {
		if val == value {
			return i, true
		}
	}
	return -1, false
}

func Last(values []int) int {
	return values[len(values)-1]
}
