package cache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	client     *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
)

func GetClient(ctx context.Context) (*s3.Client, error) {
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

	client = s3.NewFromConfig(cfg)
	return client, nil
}

func GetUploader(ctx context.Context) (*manager.Uploader, error) {
	var err error
	if client, err = GetClient(ctx); err != nil {
		return nil, err
	}

	uploader = manager.NewUploader(client)
	return uploader, nil
}

func GetDownloader(ctx context.Context) (*manager.Downloader, error) {
	var err error
	if client, err = GetClient(ctx); err != nil {
		return nil, err
	}

	downloader = manager.NewDownloader(client)
	return downloader, nil
}
