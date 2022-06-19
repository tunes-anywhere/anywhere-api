package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tunes-anywhere/anywhere-api/handler"
	"github.com/tunes-anywhere/anywhere-api/pkg/util"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{StatusCode: 200, Body: "OK"}, nil
}

func main() {
	util.InitConfig()
	lambda.Start(handler.LogHTTPRequestMiddleware(Handler))
}
