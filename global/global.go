package global

var g Global

func init() {
	g = Global{}
}

type Global struct {
	data     []float64
	fileName string
}

func DataCopy() []float64 {
	var dataCopy = make([]float64, len(g.data))
	copy(dataCopy, g.data)
	return dataCopy
}

func SetData(data []float64) {
	g.data = data
}

func FileName() string {
	return g.fileName
}

func SetFileName(fileName string) {
	g.fileName = fileName
}
