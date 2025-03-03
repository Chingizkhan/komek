package money

import "math"

// ToInt из тенге в тиины
func ToInt(f float64) int {
	return int(math.RoundToEven(f * 100))
}

// ToFloat из тиины в тенге
func ToFloat(i int64) float64 {
	if i == 0 {
		return 0
	}
	return float64(i) / 100
}
