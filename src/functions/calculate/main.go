package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/calculator"
)

type Calculate struct {
	Values    []float64            `json:"values"`
	Operation calculator.Operation `json:"operation"`
}

func publishResult(result float64) {
	log.Printf("result: %.2f", result)
}

func handler(ctx context.Context, event json.RawMessage) error {
	var calculate Calculate
	if err := json.Unmarshal(event, &calculate); err != nil {
		return err
	}

	operation, err := calculator.Initialize(calculate.Operation)
	if err != nil {
		return err
	}

	err = operation.SetValues(calculate.Values)
	if err != nil {
		return err
	}

	result := operation.Calculate()
	publishResult(result)

	return nil
}

func main() {
	lambda.Start(handler)
}
