package math

type FloatRangeMapper struct {
	Source, Target FloatRange
}

func (fmr FloatRangeMapper) Map(value float64) float64 {
	return mapUsingRanges(fmr.Source, fmr.Target, value)
}

func mapUsingRanges(source, target FloatRange, value float64) float64 {
	return (value - source.From) * target.Range() / source.Range() + target.From
}
