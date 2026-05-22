package calculator

type AddHandler struct {
	Values []float64 `json:"values"`
}

func (o *AddHandler) SetValues(values []float64) error {
	if len(values) < 2 {
		return ErrAtLeastTwoValues
	}

	o.Values = values

	return nil
}

func (o *AddHandler) Calculate() float64 {
	sum := 0.0
	for _, value := range o.Values {
		sum += value
	}
	return sum
}
