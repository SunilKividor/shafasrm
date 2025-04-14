package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GenerateNewAWSObjectKey(prefix string, id uuid.UUID, contentType string) string {
	key := uuid.New()
	var extension string

	parts := strings.Split(contentType, "/")
	if len(parts) > 1 {
		extension = parts[len(parts)-1]
	} else {
		extension = "jpg"
	}

	awsS3BucketKey := fmt.Sprintf("%s/%s/%s.%s", prefix, id.String(), key, extension)

	return awsS3BucketKey
}
