package utils

import "math"

func IsPrime(no int) bool {
	for i := 2; i <= (no - 1); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime2(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime3(no int) bool {
	end := int(math.Sqrt(float64(no)))
	for i := 2; i <= end; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
