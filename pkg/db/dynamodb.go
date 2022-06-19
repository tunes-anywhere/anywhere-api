package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	client *dynamodb.Client
)

func GetClient(ctx context.Context) (*dynamodb.Client, error) {
	if client != nil {
		return client, nil
	}

	var (
		err error
		cfg aws.Config
	)

	if cfg, err = config.LoadDefaultConfig(ctx); err != nil {
		return nil, err
	}

	client = dynamodb.NewFromConfig(cfg)
	return client, nil
}
