package calculator

import "github.com/rdzPedraos/AutomatedCanaryAnalisis/app/shared/response"

var (
	ErrDivisionByZero = response.InvalidRequestError("division by zero")
)

type DivideHandler struct {
	dividend float64
	divisor  float64
}

func (o *DivideHandler) SetValues(values []float64) error {
	if len(values) != 2 {
		return ErrMandatoryOnlyTwoValues
	}

	if values[0] == 0 {
		return ErrDivisionByZero
	}

	o.dividend = values[0]
	o.divisor = values[1]

	return nil
}

func (o *DivideHandler) Calculate() float64 {
	return o.dividend / o.divisor
}
