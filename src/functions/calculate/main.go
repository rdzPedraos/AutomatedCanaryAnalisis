package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/calculator"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/lambda"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/logger"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/response"
)

var (
	errUnmarshalRequestBody = response.InvalidRequestError("error unmarshalling request body")
)

type Calculate struct {
	Values    []float64            `json:"values"`
	Operation calculator.Operation `json:"operation"`
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var calculate Calculate

	if err := json.Unmarshal([]byte(request.Body), &calculate); err != nil {
		logger.Error(ctx, "error_unmarshalling_request_body", err)

		return response.Error(errUnmarshalRequestBody)
	}

	operation, err := calculator.Initialize(calculate.Operation)
	if err != nil {
		logger.Error(ctx, "error_initializing_operation", err)

		return response.Error(err)
	}

	err = operation.SetValues(calculate.Values)
	if err != nil {
		logger.Error(ctx, "error_setting_values", err)

		return response.Error(err)
	}

	result := operation.Calculate()

	logger.Info(ctx, "calculated_successfully", logger.Obj("result", map[string]any{
		"value": result,
	}))

	return response.Success(result)
}

func main() {
	lambda.Start(handler)
}
