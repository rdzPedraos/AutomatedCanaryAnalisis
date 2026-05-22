package calculator

import "errors"

var (
	ErrInvalidOperation       = errors.New("invalid operation")
	ErrAtLeastTwoValues       = errors.New("at least two values are required for operation")
	ErrMandatoryOnlyTwoValues = errors.New("expected only two values are required for operation")
)
