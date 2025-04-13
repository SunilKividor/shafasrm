package aws

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SunilKividor/shafasrm/internal/configs"
	"github.com/SunilKividor/shafasrm/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func GetPreSignedUrl(id uuid.UUID, presignedUrlReq models.PreSignedUrlReq) (string, error) {
	//user config
	awsUserId := os.Getenv("IAMUSERACCESSKEY")
	awsUserSecret := os.Getenv("IAMUSERSECRET")
	userCongif := configs.NewAwsUserConfig(&awsUserId, &awsUserSecret)

	log.Println(userCongif)

	//s3config
	awsS3BucketName := os.Getenv("S3BUCKETNAME")
	awsS3BucketKey := fmt.Sprintf("%s/%s", id.String(), presignedUrlReq.FileName)
	awsS3BucketRegion := os.Getenv("S3BUCKETREGION")
	awsS3Config := configs.NewAwsS3Config(&awsS3BucketName, &awsS3BucketRegion)

	log.Println(awsS3Config)

	//awsConfig
	awsConfig, err := configs.GetAwsConfig(userCongif, awsS3Config)
	if err != nil {
		return "", err
	}

	s3Client := s3.NewFromConfig(awsConfig)

	//presign url
	preSignClient := s3.NewPresignClient(s3Client)
	ctx := context.Background()
	preSignReq, err := preSignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      awsS3Config.BucketName,
		Key:         &awsS3BucketKey,
		ContentType: &presignedUrlReq.MIME,
	}, s3.WithPresignExpires(5*time.Minute),
	)

	if err != nil {
		return "", err
	}

	return preSignReq.URL, nil
}
