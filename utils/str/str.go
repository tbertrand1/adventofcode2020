package str

import "strconv"

func Atoi(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
