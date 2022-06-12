package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

type (
	HTTPHandlerFn func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error)
	AuthHandlerFn func(context.Context, events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error)
)

func LogHTTPRequestMiddleware(fn HTTPHandlerFn) HTTPHandlerFn {
	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		log.Info().
			Str("source_ip", request.RequestContext.HTTP.SourceIP).
			Str("path", request.RawPath).
			Send()

		return fn(ctx, request)
	}
}

func LogAuthRequestMiddleware(fn AuthHandlerFn) AuthHandlerFn {
	return func(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error) {
		log.Info().
			Str("source_ip", request.RequestContext.HTTP.SourceIP).
			Str("path", request.RawPath).
			Send()

		return fn(ctx, request)
	}
}
