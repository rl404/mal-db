package utils

import (
	"math"
)

// GetPercent to calculate percentage result with adjustable number of digit after comma.
// With default 2 digits after comma.
func GetPercent(top interface{}, bottom interface{}, digit ...int) float64 {
	var newTop, newBottom float64

	switch v := top.(type) {
	case int:
		newTop = float64(v)
	case float64:
		newTop = v
	}

	switch v := bottom.(type) {
	case int:
		newBottom = float64(v)
	case float64:
		newBottom = v
	}

	if newBottom == 0.0 {
		return 0.0
	}

	result := (newTop / newBottom) * 100.00

	afterComma := 2.0
	if len(digit) > 0 && digit[0] >= 0 {
		afterComma = float64(digit[0])
	}

	return math.Round(result*math.Pow(10, afterComma)) / math.Pow(10, afterComma)
}
