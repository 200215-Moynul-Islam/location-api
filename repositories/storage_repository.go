package repositories

import (
	"context"
	"time"

	"location-api/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageRepository interface {
	CopyObject(ctx context.Context, sourceKey string, destinationKey string) error
	GeneratePresignedGetObjectURL(
		ctx context.Context,
		objectKey string,
	) (string, error)
}

type storageRepository struct {
	client *s3.Client
	bucket string
}

func NewStorageRepository(client *s3.Client) StorageRepository {
	return &storageRepository{
		client: client,
		bucket: utils.GetConfig("MINIO_BUCKET"),
	}
}

func (repository *storageRepository) CopyObject(
	ctx context.Context,
	sourceKey string,
	destinationKey string,
) error {
	_, err := repository.client.CopyObject(
		ctx,
		&s3.CopyObjectInput{
			Bucket:     &repository.bucket,
			Key:        &destinationKey,
			CopySource: aws.String(repository.bucket + "/" + sourceKey),
		},
	)

	return err
}

func (repository *storageRepository) GeneratePresignedGetObjectURL(
	ctx context.Context,
	objectKey string,
) (string, error) {
	presignClient := s3.NewPresignClient(repository.client)

	expiration, err := time.ParseDuration(
		utils.GetConfig("MINIO_PRESIGNED_URL_EXPIRATION"),
	)
	if err != nil {
		return "", err
	}

	presignedRequest, err := presignClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: &repository.bucket,
			Key:    &objectKey,
		},
		func(options *s3.PresignOptions) {
			options.Expires = expiration
		},
	)
	if err != nil {
		return "", err
	}

	return presignedRequest.URL, nil
}
