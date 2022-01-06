// Package rounders provides utilities for rounding and truncating floating point numbers.
package rounders

import "math"

// RoundToDecimals rounds the given number n to the given number of decimals.
// If decimals is negative, the integer component of the number is rounded to
// abs(decimals) orders of magnitude. Decimals can be negative, which means the number
// is rounded to some integer power of 10.
func RoundToDecimals(n float64, decimals int) float64 {
	scale := math.Pow10(decimals)
	return math.Round(n*scale) / scale
}

// RoundToDecimals truncates the given number n to the given number of decimals.
// If decimals is negative, the integer component of the number is truncated to
// abs(decimals) orders of magnitude. Decimals can be negative, which means the number
// is truncated to some integer power of 10.
func TruncToDecimals(n float64, decimals int) float64 {
	scale := math.Pow10(decimals)
	return math.Trunc(n*scale) / scale
}

func toSigFigs(n float64, sigFigs int, rounding func(float64, int) float64) float64 {
	if n == 0 {
		return 0
	}

	decimalsToRound := sigFigs

	p := n
	for math.Abs(p) >= 1 {
		p /= 10
		decimalsToRound -= 1
	}

	for math.Abs(p) < 0.1 {
		p *= 10
		decimalsToRound += 1
	}

	return rounding(n, decimalsToRound)
}

// RoundToSigFigs rounds the given number n to the given number of significant figures.
// Panics if sigFigs is less than 1.
func RoundToSigFigs(n float64, sigFigs int) float64 {
	if sigFigs < 1 {
		panic("cannot round to sigfig count less than 1")
	}

	return toSigFigs(n, sigFigs, RoundToDecimals)
}

// TruncToSigFigs truncates the given number n to the given number of significant figures.
// Panics if sigFigs is less than 1.
func TruncToSigFigs(n float64, sigFigs int) float64 {
	if sigFigs < 1 {
		panic("cannot truncate to sigfig count less than 1")
	}

	return toSigFigs(n, sigFigs, TruncToDecimals)
}
