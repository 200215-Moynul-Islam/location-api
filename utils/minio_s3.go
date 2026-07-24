package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewMinIOS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(GetConfig("MINIO_REGION")),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				GetConfig("MINIO_ACCESS_KEY"),
				GetConfig("MINIO_SECRET_KEY"),
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.BaseEndpoint = aws.String(GetConfig("MINIO_ENDPOINT"))
		options.UsePathStyle = true
	}), nil
}
