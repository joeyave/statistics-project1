package templates

type Main struct {
	Title string
}

type Upload struct {
	Title string
	Data  []float64
}

type VariationalSeries struct {
	Title    string
	Image    string
	Variants []*Variant
}

type Variant struct {
	X float64
	N int
	P float64
	F float64
}

type Classes struct {
	Title   string
	Image   string
	Classes []*Class

	M int
	H float64
}

type Class struct {
	XFrom float64
	XTo   float64
	N     int
	P     float64
	F     float64
}

type Characteristics struct {
	Title           string
	Characteristics []*Characteristic
	Image           string
}

type Characteristic struct {
	Name   string
	Val    float64
	StdDev float64
	From   float64
	To     float64
}
