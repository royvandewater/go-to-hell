package challenge1

import "math"

const phi = -1 / math.Phi

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	floatN := float64(n)
	numerator := math.Pow(math.Phi, floatN) - math.Pow(phi, floatN)
	denominator := math.Sqrt(5)
	return uint64(numerator / denominator)
}
