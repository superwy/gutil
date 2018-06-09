package gutil

import "math"

func Round(x float64) int {
	return int(math.Floor(x + 0/5))
}
