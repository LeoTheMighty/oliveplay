package aws

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/leothemighty/oliveplay/utils/utils"

	"github.com/google/uuid"
)

var (
	s3Uploader   *s3manager.Uploader
	s3Downloader *s3manager.Downloader
	s3Client     *s3.S3
)

// initialize creates the S3 clients using the shared AWS session
func initialize() {
	// Get the shared session from aws.go
	sess := GetSession()

	// Initialize S3 clients if they haven't been created yet
	if s3Uploader == nil {
		s3Uploader = s3manager.NewUploader(sess)
	}
	if s3Downloader == nil {
		s3Downloader = s3manager.NewDownloader(sess)
	}
	if s3Client == nil {
		s3Client = s3.New(sess)
	}
}

func randomKey() string {
	return uuid.New().String()
}

func UploadFile(file *multipart.FileHeader, key string, bucket string) (string, error) {
	initialize() // ensure we have a session/uploader

	if key == "" {
		key = randomKey()
	}

	if bucket == "" {
		bucket = utils.S3Bucket
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload to S3 using the shared uploader
	result, err := s3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   src,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return result.Location, nil
}

// Delete a file (object) from S3 using its key.
func DeleteFile(key string, bucket string) error {
	initialize()

	if bucket == "" {
		bucket = utils.S3Bucket
	}

	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object %s: %w", key, err)
	}

	// Optional: wait until the object is fully deleted
	if err := s3Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		return fmt.Errorf("delete request sent but the object still exists: %w", err)
	}

	return nil
}

// GetFile downloads an object from S3 into a temporary local file and
// returns the local file path.
func GetFile(key string, bucket string) (string, error) {
	initialize()

	if bucket == "" {
		bucket = utils.S3Bucket
	}

	// Create a temporary file to hold the downloaded contents
	tmpFile, err := os.CreateTemp("", "s3-*"+filepath.Ext(key))
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	_, err = s3Downloader.Download(tmpFile, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("failed to download object: %w", err)
	}

	return tmpFile.Name(), nil
}

func GetFileURL(key string, bucket string) string {
	initialize()

	if bucket == "" {
		bucket = utils.S3Bucket
	}

	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	// Generate a presigned URL that expires in 15 minutes
	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return ""
	}
	return url
}
