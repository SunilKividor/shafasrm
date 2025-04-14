package util

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateNewAWSObjectKey(prefix string, id uuid.UUID) string {
	key := uuid.New()
	awsS3BucketKey := fmt.Sprintf("%s/%s/%s", prefix, id.String(), key)

	return awsS3BucketKey
}

// func GenerateAWSObjectKey(prefix string, id uuid.UUID) string {
// 	key := uuid.New()
// 	awsS3BucketKey := fmt.Sprintf("%s/%s/%s", prefix, id.String(), key)

// 	return awsS3BucketKey
// }
