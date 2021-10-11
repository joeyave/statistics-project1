package helpers

import (
	"github.com/joeyave/statistics-project1/templates"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"sort"
)

func Variance(x []float64) float64 {
	sum := 0.
	mean := Mean(x)
	for _, val := range x {
		sum += math.Pow(val-mean, 2)
	}

	return sum / float64(len(x)-1)
}

func VarianceShifted(x []float64) float64 {
	sum := 0.
	mean := Mean(x)
	for _, val := range x {
		sum += math.Pow(val-mean, 2)
	}

	return sum / float64(len(x))
}

func StandardDeviation(x []float64) float64 {
	variance := Variance(x)
	stdDev := math.Sqrt(variance)
	return stdDev
}

func StandardDeviationShifted(x []float64) float64 {
	variance := VarianceShifted(x)
	stdDev := math.Sqrt(variance)
	return stdDev
}

func Mean(x []float64) float64 {
	n := float64(len(x))

	sum := 0.0
	for _, v := range x {
		sum += v
	}
	mean := sum / n

	return mean
}

func MeanStandardError(x []float64) float64 {
	stdDev := StandardDeviation(x)
	stdErr := stdDev / math.Sqrt(float64(len(x)))
	return stdErr
}

func MeanConfidenceInterval(alpha float64, x []float64) (float64, float64) {
	mean := Mean(x)
	stdErr := MeanStandardError(x)
	v := float64(len(x) - 1)
	t := QuantileT(1-alpha/2, v)

	low := mean - t*stdErr
	high := mean + t*stdErr

	return low, high
}

func Median(x []float64) float64 {
	n := len(x)

	if n == 1 {
		return x[0]
	}

	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}

	med := 0.
	if n%2 == 0 {
		med = (x[n/2+1]-x[n/2])/2 + x[n/2]
	} else {
		med = x[n/2]
	}

	return med
}

func MedianConfidenceInterval(alpha float64, x []float64) (float64, float64) {
	// https://www-users.york.ac.uk/~mb55/intro/cicent.htm

	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}

	u := QuantileU(1 - alpha/2)

	i := int(math.Floor(float64(len(x))/2 - u*math.Sqrt(float64(len(x)))/2))
	k := int(math.Floor(float64(len(x))/2 + 1 + u*math.Sqrt(float64(len(x)))/2))

	if k > len(x)-1 {
		k = len(x) - 1
	}

	return x[i], x[k]
}

func StandardDeviationStandardError(x []float64) float64 {
	stdDev := StandardDeviation(x)
	stdDevStdErr := stdDev / math.Sqrt(float64(2*len(x)))
	return stdDevStdErr
}

func StandardDeviationConfidenceInterval(alpha float64, x []float64) (float64, float64) {

	stdDev := StandardDeviation(x)
	stdDevOfStdDev := stdDev / math.Sqrt(float64(2*len(x)))

	v := float64(len(x) - 1)
	t := QuantileT(1-alpha/2, v)

	return stdDev - t*stdDevOfStdDev, stdDev + t*stdDevOfStdDev
}

func Skewness(x []float64) float64 {
	sum := 0.
	mean := Mean(x)
	for _, val := range x {
		sum += math.Pow(val-mean, 3)
	}

	stdDev := StandardDeviationShifted(x)

	skewness := sum / (float64(len(x)) * math.Pow(stdDev, 3))
	return skewness
}

func SkewnessStandardError(x []float64) float64 {
	N := len(x)
	stdErr := math.Sqrt(float64(6*(N-2)) / float64((N+1)*(N+3)))
	return stdErr
}

func SkewnessConfidenceInterval(alpha float64, x []float64) (float64, float64) {
	skewness := Skewness(x)
	t := QuantileT(1-alpha/2, float64(len(x)-1))
	skewnessStdErr := SkewnessStandardError(x)

	low := skewness - t*skewnessStdErr
	high := skewnessStdErr + t*skewnessStdErr

	return low, high
}

func Kurtosis(x []float64) float64 {
	sum := 0.
	mean := Mean(x)
	for _, val := range x {
		sum += math.Pow(val-mean, 4)
	}

	stdDev := StandardDeviationShifted(x)

	kurtosis := (sum / (float64(len(x)) * math.Pow(stdDev, 4))) - 3
	return kurtosis
}

func KurtosisStandardError(x []float64) float64 {
	N := len(x)
	stdErr := math.Sqrt(float64(24*N*(N-2)*(N-3)) / (math.Pow(float64(N+1), 2) * float64((N+3)*(N+5))))
	return stdErr
}

func KurtosisConfidenceInterval(alpha float64, x []float64) (float64, float64) {
	kurtosis := Kurtosis(x)
	t := QuantileT(1-alpha/2, float64(len(x)-1))
	kurtosisStdErr := KurtosisStandardError(x)

	low := kurtosis - t*kurtosisStdErr
	high := kurtosisStdErr + t*kurtosisStdErr

	return low, high
}

func AntiKurtosis(x []float64) float64 {
	kurtosis := Kurtosis(x)
	antiKurtosis := 1 / math.Sqrt(kurtosis+3)
	return antiKurtosis
}

func ECDF(x []float64) []*templates.Variant {

	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}

	variantToNumMap := make(map[float64]int)

	for _, v := range x {
		_, exists := variantToNumMap[v]
		if exists {
			variantToNumMap[v] += 1
		} else {
			variantToNumMap[v] = 1
		}
	}

	keys := make([]float64, 0, len(variantToNumMap))
	for k := range variantToNumMap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	var variants []*templates.Variant

	for _, v := range keys {
		num := variantToNumMap[v]
		variant := templates.Variant{
			X: v,
			N: num,
			P: float64(num) / float64(len(x)),
		}

		variants = append(variants, &variant)
	}

	for i, v := range variants {

		if i == 0 {
			v.F = v.P
		} else {
			v.F = variants[i-1].F + v.P
		}
	}

	return variants
}

func PlotECDF(x []float64) *plot.Plot {

	variants := ECDF(x)

	p := plot.New()
	p.Add(plotter.NewGrid())
	p.X.Label.Text = "x"
	p.Y.Label.Text = "ECDF"

	p.Y.Min = 0
	p.X.Min = 0

	from := plotter.NewFunction(func(x_i float64) float64 {
		if x_i < x[0] {
			return 0
		}
		return math.NaN()
	})
	from.Color = color.RGBA{R: 255, A: 255}

	p.Add(from)

	for i := 0; i < len(variants)-1; i++ {
		dot1 := plotter.XY{X: variants[i].X, Y: variants[i].F}

		dot2 := plotter.XY{}
		if i == len(variants)-1 {
			dot2 = plotter.XY{X: variants[len(variants)-1].X + 1, Y: 1}
		} else {
			dot2 = plotter.XY{X: variants[i+1].X, Y: variants[i].F}
		}

		err := plotutil.AddLines(p, plotter.XYs{dot1, dot2})
		if err != nil {
			return nil
		}

		err = plotutil.AddScatters(p, plotter.XYs{dot1})
		if err != nil {
			return nil
		}
	}

	return p
}

func Classes(M int, x []float64) []*templates.Class {

	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}

	var classes []*templates.Class

	xMin := Min(x)
	xMax := Max(x)

	h := (xMax - xMin) / float64(M)

	for i := 0; i < M; i++ {
		class := templates.Class{
			XFrom: xMin + (h * float64(i)),
		}
		class.XTo = class.XFrom + h

		classes = append(classes, &class)
	}

	for i := range classes {
		for _, v := range x {
			if i == len(classes)-1 {
				if v >= classes[i].XFrom && v <= classes[i].XTo {
					classes[i].N++
				}
			} else {
				if v >= classes[i].XFrom && v < classes[i].XTo {
					classes[i].N++
				}
			}
		}
	}

	for i := range classes {
		classes[i].P = float64(classes[i].N) / float64(len(x))

		if i == 0 {
			classes[i].F = classes[i].P
		} else {
			classes[i].F = classes[i-1].F + classes[i].P
		}
	}

	return classes
}

func Scott(x []float64) float64 {
	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}
	stdDev := StandardDeviation(x)
	return stdDev * math.Pow(float64(len(x)), -0.2)
}

func KDE(h float64, x []float64) func(x float64) float64 {
	return func(x_i float64) float64 {
		if !sort.Float64sAreSorted(x) {
			sort.Float64s(x)
		}

		kSum := 0.
		for _, val := range x {
			u := (x_i - val) / h
			k := 1 / math.Sqrt(2*math.Pi) * math.Exp(-(math.Pow(u, 2) / 2))
			kSum += k
		}

		y := (1 / (float64(len(x)) * h)) * kSum

		return y
	}
}

func PlotHistogram(M int, h float64, x []float64) *plot.Plot {

	variants := ECDF(x)

	p := plot.New()
	p.X.Label.Text = "x"
	p.Y.Label.Text = "p"

	yMax := 0.
	for _, v := range x {
		y := KDE(h, x)(v)
		if y > yMax {
			yMax = y
		}
	}
	p.Y.Max = yMax + 0.01

	var XYs plotter.XYs

	for _, v := range variants {
		xy := plotter.XY{X: v.X, Y: v.P}
		XYs = append(XYs, xy)
	}

	histogram, err := plotter.NewHistogram(XYs, M)
	if err != nil {
		return nil
	}

	p.Add(histogram)

	kde := plotter.NewFunction(KDE(h, x))

	kde.Color = color.RGBA{R: 255, A: 255}
	kde.Width = vg.Points(2)
	p.Add(kde)

	return p
}

func PDF(x []float64) func(x_i float64) float64 {
	return func(x_i float64) float64 {
		stdDev := StandardDeviation(x)
		mean := Mean(x)
		y := math.Pow(math.E, -0.5*math.Pow((x_i-mean)/stdDev, 2)) / stdDev * math.Sqrt(2*math.Pi)
		return y
	}
}

func PlotPDF(x []float64) *plot.Plot {
	// https://en.wikipedia.org/wiki/Normal_distribution

	p := plot.New()
	p.Add(plotter.NewGrid())

	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	yMax := 0.
	for _, v := range x {
		y := PDF(x)(v)
		if y > yMax {
			yMax = y
		}
	}

	p.X.Min = Min(x)
	p.X.Max = Max(x)

	p.Y.Min = 0
	p.Y.Max = yMax + yMax*0.1

	pdf := plotter.NewFunction(PDF(x))

	p.Add(pdf)

	return p
}

func Min(x []float64) float64 {
	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}
	return x[0]
}

func Max(x []float64) float64 {
	if !sort.Float64sAreSorted(x) {
		sort.Float64s(x)
	}
	return x[len(x)-1]
}

func QuantileU(p float64) float64 {

	phi := func(a float64) float64 {
		const c0, c1, c2, d1, d2, d3 = 2.515517, 0.802853, 0.010328, 1.432788, 0.1892659, 0.001308

		t := math.Sqrt(-2 * math.Log(a))

		return t - ((c0 + c1*t + c2*math.Pow(t, 2)) / (1 + d1*t + d2*math.Pow(t, 2) + d3*math.Pow(t, 3)))
	}

	if p <= 0.5 {
		return -phi(p)
	}
	return phi(1 - p)
}

func QuantileT(p, v float64) float64 {
	u := QuantileU(p)

	g1 := (math.Pow(u, 3) + u) / 4
	g2 := (5*math.Pow(u, 5) + 16*math.Pow(u, 3) + 3*u) / 96
	g3 := (3*math.Pow(u, 7) + 19*math.Pow(u, 5) + 17*math.Pow(u, 3) - 15*u) / 384
	g4 := (79*math.Pow(u, 9) + 779*math.Pow(u, 7) + 1482*math.Pow(u, 5) - 1920*math.Pow(u, 3) - 945*u) / 92160

	return u + g1/v + g2/math.Pow(v, 2) + g3/math.Pow(v, 3) + g4/math.Pow(v, 4)
}
