package lambda

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	root "github.com/aws/aws-lambda-go/lambda"
	"github.com/rdzPedraos/AutomatedCanaryAnalisis/src/libraries/logger"
)

type Handler func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

func Start(handler Handler) {
	root.Start(func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		startTime := time.Now()

		response, err := handler(ctx, request)

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		logger.Auto(ctx, "request_finished", err, logger.Obj("metrics", map[string]any{
			"duration":    duration,
			"status_code": response.StatusCode,
		}))

		return response, err
	})
}
