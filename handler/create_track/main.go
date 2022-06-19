package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"
	"github.com/tunes-anywhere/anywhere-api/handler"
	"github.com/tunes-anywhere/anywhere-api/pkg/models/track"
	"github.com/tunes-anywhere/anywhere-api/pkg/util"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if request.Body == "" {
		return handler.HTTPResponse{}.
			BadRequestf("must provide a body").
			Response()
	}

	var (
		err error
		raw track.TrackAttributes
	)

	if err = json.Unmarshal([]byte(request.Body), &raw); err != nil {
		return handler.HTTPResponse{}.
			BadRequestf("failed to parse body: %s", err.Error()).
			Response()
	}

	log.Debug().Interface("raw", raw).Send()
	// TODO save track

	return handler.HTTPResponse{}.
		Ok("OK", nil).
		Response()
}

func main() {
	util.InitConfig()
	lambda.Start(handler.LogHTTPRequestMiddleware(Handler))
}
