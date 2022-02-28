package utils

import "math"

func RoundUp(value float32) float32 {
	return float32(math.Ceil(float64(value)*100) / 100)
}