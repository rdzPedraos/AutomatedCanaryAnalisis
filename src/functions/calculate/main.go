package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/calculator"
)

type Calculate struct {
	Values    []float64            `json:"values"`
	Operation calculator.Operation `json:"operation"`
}

func responseError(statusCode int, err error) (events.APIGatewayV2HTTPResponse, error) {
	resp, _ := response(statusCode, err.Error())
	return resp, err
}

func response(statusCode int, body string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "application/text"},
	}, nil
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var calculate Calculate
	if err := json.Unmarshal([]byte(request.Body), &calculate); err != nil {
		return responseError(400, err)
	}

	operation, err := calculator.Initialize(calculate.Operation)
	if err != nil {
		return responseError(400, err)
	}

	err = operation.SetValues(calculate.Values)
	if err != nil {
		return responseError(400, err)
	}

	result := operation.Calculate()

	return response(200, fmt.Sprintf("%.2f", result))
}

func main() {
	lambda.Start(handler)
}
