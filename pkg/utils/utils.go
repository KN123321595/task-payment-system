package utils

import "math/rand"

func GenerateRandomSequence(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}

	return string(result)
}
