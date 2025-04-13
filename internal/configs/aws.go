package configs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type AwsUserConfig struct {
	ID     *string
	Secret *string
}

func NewAwsUserConfig(id *string, secret *string) *AwsUserConfig {
	return &AwsUserConfig{
		ID:     id,
		Secret: secret,
	}
}

type AwsS3Config struct {
	BucketName *string
	Region     *string
}

func NewAwsS3Config(bucketName *string, region *string) *AwsS3Config {
	return &AwsS3Config{
		BucketName: bucketName,
		Region:     region,
	}
}

func GetAwsConfig(awsUser *AwsUserConfig, awsS3 *AwsS3Config) (aws.Config, error) {

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(*awsS3.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			*awsUser.ID,
			*awsUser.Secret,
			"",
		)),
	)

	return cfg, err
}
