package templates

import "github.com/shopspring/decimal"

type Main struct {
	Title string
}

type Upload struct {
	Title string
	Data  []decimal.Decimal
}

type VariationalSeries struct {
	Title    string
	Variants []*Variant
}

type Variant struct {
	X decimal.Decimal
	N int
	P decimal.Decimal
	F decimal.Decimal
}

type Classes struct {
	Title   string
	Classes []*Class
}

type Class struct {
	XFrom decimal.Decimal
	XTo   decimal.Decimal
	N     int
	P     decimal.Decimal
	F     decimal.Decimal
}
