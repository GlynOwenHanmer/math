package math

type FloatRangeMapper struct {
	Source, Target FloatRange
}

func (fmr FloatRangeMapper) Map(value float64) float64 {
	return (value - fmr.Source.From) * fmr.Target.Range() / fmr.Source.Range() + fmr.Target.From
}
