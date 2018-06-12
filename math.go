package gutil

func Round(x float64) int {
	if x < 0.0 {
		x -= 0.5
	} else {
		x += 0.5
	}
	return int(x)
}
