package calculator

type MultiplyHandler struct {
	Values []float64 `json:"values"`
}

func (o *MultiplyHandler) SetValues(values []float64) error {
	if len(values) < 2 {
		return ErrAtLeastTwoValues
	}

	o.Values = values

	return nil
}

func (o *MultiplyHandler) Calculate() float64 {
	multiply := o.Values[0]
	for _, value := range o.Values[1:] {
		multiply *= value
	}
	return multiply
}
