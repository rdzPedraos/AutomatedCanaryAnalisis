package calculator

type SubstractHandler struct {
	Values []float64 `json:"values"`
}

func (o *SubstractHandler) SetValues(values []float64) error {
	if len(values) < 2 {
		return ErrAtLeastTwoValues
	}

	o.Values = values

	return nil
}

func (o *SubstractHandler) Calculate() float64 {
	subtract := o.Values[0]
	for _, value := range o.Values[1:] {
		subtract -= value
	}
	return subtract
}
