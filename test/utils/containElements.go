package utils

func ContainElements[T Int | Float | UInt](arr1, arr2 []T) bool {
	var numbersMap = make(map[T]int)

	for _, number := range arr1 {
		numbersMap[number]++
	}

	for _, number2 := range arr2 {
		if count, ok := numbersMap[number2]; !ok || count == 0 {
			return false
		}
		numbersMap[number2]--
	}

	return true
}
