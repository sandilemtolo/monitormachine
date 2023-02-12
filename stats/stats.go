package stats

type Reading struct {
	measure string
	unit    string
}

type Statistics interface {
	Collect() Reading
}
