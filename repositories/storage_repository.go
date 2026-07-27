package repositories

import (
	"context"

	"location-api/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageRepository interface {
	GetObject(ctx context.Context, objectKey string) (*s3.GetObjectOutput, error)
	CopyObject(ctx context.Context, sourceKey string, destinationKey string) error
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

func (repository *storageRepository) GetObject(
	ctx context.Context,
	objectKey string,
) (*s3.GetObjectOutput, error) {
	return repository.client.GetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: &repository.bucket,
			Key:    &objectKey,
		},
	)
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
