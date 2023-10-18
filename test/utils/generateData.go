package utils

import "math/rand"

func GenerateRandomData(len int) []float32 {
	arr := make([]float32, len)

	for i := 0; i < len; i++ {
		arr[i] = rand.Float32()
	}

	return arr
}
