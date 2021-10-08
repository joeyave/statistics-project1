package helpers

import (
	"gonum.org/v1/gonum/stat"
	"math"
	"sort"
)

func Mean(x []float64) float64 {
	n := float64(len(x))

	sum := 0.0
	for _, v := range x {
		sum += v
	}
	mean := sum / n
	return mean
}

func Median(x []float64) float64 {
	n := len(x)
	if n == 0 {
		// TODO: return an error
		return 0
	}
	if n == 1 {
		return x[0]
	}

	sort.Float64s(x)

	med := 0.
	if n%2 == 0 {
		med = (x[n/2+1]-x[n/2])/2 + x[n/2]
	} else {
		med = x[n/2]
	}

	return med
}

func Kurtosis(x []float64) float64 {
	mean := Mean(x)

	sum := 0.
	for _, val := range x {
		sum += math.Pow(val-mean, 4)
	}

	S := stat.Variance(x, nil)
	kurtosis := (1 / (float64(len(x)) * math.Pow(S, 2)) * sum) - 3
	return kurtosis
}

func AntiKurtosis(x []float64) float64 {
	kurtosis := Kurtosis(x)
	antiKurtosis := 1 / math.Sqrt(kurtosis+3)
	return antiKurtosis
}
