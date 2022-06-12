package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hans-m-song/anywhere/handler"
	"github.com/hans-m-song/anywhere/pkg/util"
)

func Handler(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error) {
	resourceARNPrefix := fmt.Sprintf(
		"arn:aws:execute-api:*:%s:%s/%s",
		request.RequestContext.AccountID,
		request.RequestContext.APIID,
		request.RequestContext.Stage,
	)

	// TODO narrow for different users
	resources := []string{fmt.Sprintf("%s/%s", resourceARNPrefix, "*")}

	return events.APIGatewayV2CustomAuthorizerIAMPolicyResponse{
		PrincipalID: "Anywhere Authorization",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: resources,
				},
			},
		},
	}, nil
}

func main() {
	util.InitConfig()
	lambda.Start(handler.LogAuthRequestMiddleware(Handler))
}
