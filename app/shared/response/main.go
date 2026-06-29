package response

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func buildResponse(statusCode int, stringBody string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Body:       stringBody,
		Headers:    map[string]string{"Content-Type": "application/text"},
	}
}
func Error(err error) (events.APIGatewayV2HTTPResponse, error) {
	if apiError, ok := err.(ApiError); ok {
		return buildResponse(apiError.StatusCode, apiError.Error()), err
	}

	return buildResponse(http.StatusInternalServerError, "internal server error"), err
}

func Success(body any) (events.APIGatewayV2HTTPResponse, error) {
	return buildResponse(http.StatusOK, fmt.Sprintf("%v", body)), nil
}
