package main

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func GenerateFileName(filename string) string {
	t := time.Now()
	_ = t.String()
	currentDatetime := t.Format("02-01-2006-15-04-05PM")
	uid := strings.ToUpper(uuid.New().String()[0:6])
	newfilename := currentDatetime + "-" + uid + "-" + filename
	return newfilename
}

func UploadBulkFile(files []*multipart.FileHeader) (*[]FileUploadResponse, error) {
	conf := NewConfig(".env")
	if conf == nil {
		log.Println("cannot get configuration s3 bukcet")
		return nil, errors.New("cannot get configuration s3 bucket")
	}
	accessKey := conf.AwsAccessKeyId
	accessSecret := conf.AwsSecretAccessKey
	defaultRegion := conf.AwsDefaultRegion
	bucketName := conf.AwsStorageBucketName
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
		log.Println("aws configuration missing")
		return nil, errors.New("aws configuration missing")
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String(defaultRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // The minimum/default allowed part size is 10MB
		u.Concurrency = 5             // default is 5
	})
	entitys := make([]FileUploadResponse, 0)
	for _, file := range files {
		log.Println("Start to upload filename: ", file.Filename)
		extension := filepath.Ext(file.Filename)
		fileName := file.Filename[:len(file.Filename)-len(extension)]
		src, err := file.Open()
		if err != nil {
			log.Printf("failed to open file: %v", err)
			entitys = append(entitys, FileUploadResponse{
				Name:         fileName,
				Url:          "",
				Size:         int(file.Size),
				Extension:    extension,
				UploadStatus: false,
				Message:      "failed to open file",
			})
			continue
		}
		defer src.Close()

		key := "files/" + GenerateFileName(file.Filename)

		// Upload the file to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
			Body:   src,
			// ACL:    aws.String("public-read"),
		})

		//in case it fails to upload
		if err != nil || result == nil {
			log.Printf("failed to upload file to s3 bucket: %v", err)
			entitys = append(entitys, FileUploadResponse{
				Name:         fileName,
				Url:          "",
				Size:         int(file.Size),
				Extension:    extension,
				UploadStatus: false,
				// Message:   "failed to upload file to s3",
				Message: strings.Split(err.Error(), "\n")[0],
			})
		} else {
			url := result.Location
			log.Println("successfully file upload to s3 bucket url: ", url)
			entitys = append(entitys, FileUploadResponse{
				Name:         fileName,
				Url:          url,
				Size:         int(file.Size),
				Extension:    extension,
				UploadStatus: true,
				Message:      "upload file successfully to s3 bucket",
			})
		}
	}
	return &entitys, nil
}

func UploadSingleFile(file *multipart.FileHeader) (*FileUploadResponse, error) {
	entityResponse := &FileUploadResponse{}
	conf := NewConfig(".env")
	if conf == nil {
		log.Println("cannot get configuration s3 bukcet")
		return nil, errors.New("cannot get configuration s3 bucket")
	}
	accessKey := conf.AwsAccessKeyId
	accessSecret := conf.AwsSecretAccessKey
	defaultRegion := conf.AwsDefaultRegion
	bucketName := conf.AwsStorageBucketName
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
		log.Println("aws configuration missing")
		return nil, errors.New("aws configuration missing")
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String(defaultRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // The minimum/default allowed part size is 10MB
		u.Concurrency = 5             // default is 5
	})
	log.Println("Start to upload filename: ", file.Filename)
	extension := filepath.Ext(file.Filename)
	fileName := file.Filename[:len(file.Filename)-len(extension)]
	src, err := file.Open()
	if err != nil {
		log.Printf("failed to open file: %v", err)
		return nil, fmt.Errorf("failed to open file: %v", strings.Split(err.Error(), "\n")[0])
	}
	defer src.Close()

	key := "files/" + GenerateFileName(file.Filename)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   src,
		// ACL:    aws.String("public-read"),
	})

	//in case it fails to upload
	if err != nil || result == nil {
		log.Printf("failed to upload file to s3: %v", err)
		return nil, fmt.Errorf("failed to upload file to s3 bucket: %v", strings.Split(err.Error(), "\n")[0])
	} else {
		url := result.Location
		log.Println("successfully file upload to s3 bucket url: ", url)
		entityResponse.Name = fileName
		entityResponse.Url = url
		entityResponse.Size = int(file.Size)
		entityResponse.Extension = extension
		entityResponse.UploadStatus = true
		entityResponse.Message = "upload file successfully to s3 bucket"
	}
	return entityResponse, nil
}
