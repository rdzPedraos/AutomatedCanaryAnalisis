package calculator

type Operation string

const (
	OperationAdd      Operation = "add"
	OperationSubtract Operation = "subtract"
	OperationMultiply Operation = "multiply"
	OperationDivide   Operation = "divide"
)

type Handler interface {
	SetValues(values []float64) error
	Calculate() float64
}

var mapOperationHandler = map[Operation]Handler{
	OperationAdd:      &AddHandler{},
	OperationSubtract: &SubstractHandler{},
	OperationMultiply: &MultiplyHandler{},
	OperationDivide:   &DivideHandler{},
}

func Initialize(operation Operation) (Handler, error) {
	handler, ok := mapOperationHandler[operation]
	if !ok {
		return nil, ErrInvalidOperation
	}

	return handler, nil
}
