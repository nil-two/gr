package main

import (
	"math"
)

var goldenRatio = (1.0 + math.Sqrt(5.0)) / 2.0

func round(f float64) int {
	return int(math.Floor(f + 0.5))
}
