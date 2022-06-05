package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

type HandlerFn func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)

func LogMiddleware(fn HandlerFn) HandlerFn {
	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		log.Debug().
			Str("source_ip", request.RequestContext.HTTP.SourceIP).
			Str("path", request.RawPath).
			Send()

		return fn(ctx, request)
	}
}
