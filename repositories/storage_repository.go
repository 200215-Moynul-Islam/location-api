package repositories

import (
	"context"

	"location-api/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageRepository interface {
	ListLocationFolders(ctx context.Context) ([]string, error)
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

func (repository *storageRepository) ListLocationFolders(ctx context.Context) ([]string, error) {
	output, err := repository.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:    &repository.bucket,
		Delimiter: aws.String("/"),
	})

	if err != nil {
		return nil, err
	}

	folders := make([]string, 0, len(output.CommonPrefixes))

	for _, prefix := range output.CommonPrefixes {
		if prefix.Prefix != nil {
			folders = append(folders, *prefix.Prefix)
		}
	}

	return folders, nil
}
