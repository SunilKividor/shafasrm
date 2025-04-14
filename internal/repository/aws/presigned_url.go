package aws

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SunilKividor/shafasrm/internal/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

const (
	s3UploadKeyPrefix     = "users/uploads"
	defaultUploadExpiry   = 5 * time.Minute
	defaultDownloadExpiry = 30 * time.Minute
)

type PreSigner interface {
	GenerateUploadUrl(ctx context.Context, userID uuid.UUID, contentType string) (string, string, error)
	GenerateDownloadUrl(ctx context.Context, key string) (string, error)
	VerifyObjectExists(ctx context.Context, key string) (bool, error)
}

type PresignS3Service struct {
	preSignClient      *s3.PresignClient
	s3Client           *s3.Client
	bucketName         string
	uploadExpiry       time.Duration
	downloadExpiry     time.Duration
	uploadKeyGenerator func(userID uuid.UUID) string
}

func NewPresignS3Service(cfg aws.Config, bucketName string) (*PresignS3Service, error) {
	if bucketName == "" {
		return nil, errors.New("S3 bucket name cannot be empty")
	}

	s3Client := s3.NewFromConfig(cfg)
	preSignClient := s3.NewPresignClient(s3Client)

	defaultUploadGenerator := func(userID uuid.UUID) string {
		return util.GenerateNewAWSObjectKey(s3UploadKeyPrefix, userID)
	}

	return &PresignS3Service{
		preSignClient:      preSignClient,
		s3Client:           s3Client,
		bucketName:         bucketName,
		uploadExpiry:       defaultUploadExpiry,
		downloadExpiry:     defaultDownloadExpiry,
		uploadKeyGenerator: defaultUploadGenerator,
	}, nil
}

func (p *PresignS3Service) GenerateUploadUrl(ctx context.Context, userID uuid.UUID, contentType string) (string, string, error) {
	if contentType == "" {
		return "", "", errors.New("content type is required")
	}
	preSignClient := p.preSignClient

	key := p.uploadKeyGenerator(userID)

	preSignReq, err := preSignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      &p.bucketName,
		Key:         &key,
		ContentType: &contentType,
	}, s3.WithPresignExpires(p.uploadExpiry),
	)

	if err != nil {
		log.Printf("Failed to generate pre-signed PUT URL for key %s: %v", key, err)
		return "", "", fmt.Errorf("failed to sign upload request: %w", err)
	}

	return preSignReq.URL, key, nil
}

func (p *PresignS3Service) GenerateDownloadUrl(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("key is required")
	}
	preSignClient := p.preSignClient

	preSignReq, err := preSignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &p.bucketName,
		Key:    &key,
	}, s3.WithPresignExpires(p.downloadExpiry),
	)

	if err != nil {
		log.Printf("Failed to generate pre-signed GET URL for key %s: %v", key, err)
		return "", fmt.Errorf("failed to sign download request: %w", err)
	}

	return preSignReq.URL, nil
}

func (p *PresignS3Service) VerifyObjectExists(ctx context.Context, key string) (bool, error) {
	if key == "" {
		return false, errors.New("key is required")
	}

	_, err := p.s3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &p.bucketName,
		Key:    &key,
	})

	if err != nil {
		var nsk *types.NoSuchKey
		var nf *types.NotFound

		if errors.As(err, &nsk) || errors.As(err, &nf) {
			log.Printf("Verification : object %s not found,", key)
			return false, nil
		}

		log.Printf("Error Verifying object %s:%v", key, err)
		return false, err
	}

	log.Printf("Verification: Object %s found.", key)
	return true, nil
}
