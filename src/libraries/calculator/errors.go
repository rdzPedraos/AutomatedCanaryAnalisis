package calculator

import "github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/response"

var (
	ErrInvalidOperation       = response.InvalidRequestError("invalid operation")
	ErrAtLeastTwoValues       = response.InvalidRequestError("at least two values are required for operation")
	ErrMandatoryOnlyTwoValues = response.InvalidRequestError("expected only two values are required for operation")
)
