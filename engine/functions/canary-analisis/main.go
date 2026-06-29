package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type CanaryAnalisisRequest struct {
	Environment string `json:"environment"`
	Service     string `json:"service"`
	Version     string `json:"version"`
}

func handler(ctx context.Context, request CanaryAnalisisRequest) error {
	fmt.Println("canary_analisis_started")

	return nil
}

func main() {
	lambda.Start(handler)
}
