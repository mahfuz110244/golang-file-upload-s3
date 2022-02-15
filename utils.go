package main

import (
	"errors"
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

// func GenerateFileName(filename string) string {
// 	t := time.Now()
// 	_ = t.String()
// 	currentDatetime := t.Format("02-01-2006-15-04-05PM")
// 	uid := strings.ToUpper(uuid.New().String()[0:6])
// 	newfilename := currentDatetime + "-" + uid + "-" + filename
// 	return newfilename
// }

// func UploadBulkImage(files []*multipart.FileHeader) (*[]FileUploadResponse, error) {
// 	conf := NewConfig(".env")
// 	if conf == nil {
// 		log.Println("cannot get configuration s3 bukcet")
// 		return nil, errors.New("custom_error!!!cannot get configuration s3 bucket")
// 	}
// 	accessKey := conf.AwsAccessKeyId
// 	accessSecret := conf.AwsSecretAccessKey
// 	defaultRegion := conf.AwsDefaultRegion
// 	bucketName := conf.AwsStorageBucketName
// 	var awsConfig *aws.Config
// 	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
// 		log.Println("aws configuration missing")
// 		return nil, errors.New("custom_error!!!aws configuration missing")
// 	} else {
// 		awsConfig = &aws.Config{
// 			Region:      aws.String(defaultRegion),
// 			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
// 		}
// 	}

// 	// The session the S3 Uploader will use
// 	sess := session.Must(session.NewSession(awsConfig))

// 	// Create an uploader with the session and custom options
// 	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
// 		u.PartSize = 10 * 1024 * 1024 // The minimum/default allowed part size is 10MB
// 		u.Concurrency = 5             // default is 5
// 	})
// 	entitys := make([]FileUploadResponse, 0)
// 	for _, file := range files {
// 		src, err := file.Open()
// 		if err != nil {
// 			log.Printf("failed to open image: %v", err)
// 		}
// 		defer src.Close()

// 		extension := filepath.Ext(file.Filename)
// 		imageName := file.Filename[:len(file.Filename)-len(extension)]

// 		key := "files/" + GenerateFileName(file.Filename)

// 		// Upload the file to S3.
// 		result, err := uploader.Upload(&s3manager.UploadInput{
// 			Bucket: aws.String(bucketName),
// 			Key:    aws.String(key),
// 			Body:   src,
// 			// ACL:    aws.String("public-read"),
// 		})

// 		//in case it fails to upload
// 		if err != nil {
// 			log.Printf("failed to upload image to s3: %v", err)
// 			// return nil, errors.New("custom_error!!!failed to upload image to s3")
// 		}
// 		log.Println("successfully image upload to s3 bucket url: ", result.Location)
// 		entitys = append(entitys, FileUploadResponse{
// 			Name:      imageName,
// 			Url:       result.Location,
// 			Size:      int(file.Size),
// 			Extension: extension,
// 		})
// 	}
// 	return &entitys, nil
// }

// func SingleImageUpload(params *entity.ProductImagesUpload, conf *config.Config) (*entity.ProductImages, error) {
// 	accessKey := conf.Aws.AwsAccessKeyId
// 	accessSecret := conf.Aws.AwsSecretAccessKey
// 	defaultRegion := conf.Aws.AwsDefaultRegion
// 	bucketName := conf.Aws.AwsStorageBucketName
// 	var awsConfig *aws.Config
// 	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
// 		log.Println("aws configuration missing")
// 		return nil, errors.New("custom_error!!!aws configuration missing")
// 	} else {
// 		awsConfig = &aws.Config{
// 			Region:      aws.String(defaultRegion),
// 			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
// 		}
// 	}

// 	// The session the S3 Uploader will use
// 	sess := session.Must(session.NewSession(awsConfig))

// 	// Create an uploader with the session and custom options
// 	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
// 		u.PartSize = 10 * 1024 * 1024 // The minimum/default allowed part size is 10MB
// 		u.Concurrency = 5             // default is 5
// 	})

// 	imageEntity := &entity.ProductImages{}
// 	if conf == nil {
// 		log.Println("cannot get configuration s3 bukcet")
// 		return nil, errors.New("custom_error!!!cannot get configuration s3 bucket")
// 	}
// 	src, err := params.ImageFile.Open()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer src.Close()

// 	extension := filepath.Ext(params.ImageFile.Filename)
// 	imageName := params.ImageFile.Filename[:len(params.ImageFile.Filename)-len(extension)]

// 	key := "products/images/" + GenerateFileName(params.ImageFile.Filename, params.ProductID)

// 	// Upload the file to S3.
// 	result, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(key),
// 		Body:   src,
// 		// ACL:    aws.String("public-read"),
// 	})

// 	//in case it fails to upload
// 	if err != nil {
// 		log.Printf("failed to upload image to s3: %v", err)
// 		return nil, errors.New("custom_error!!!failed to upload image to s3")
// 	}
// 	log.Println("successfully image upload to s3 bucket url: ", result.Location)
// 	imageEntity.ImageUrl = result.Location
// 	imageEntity.ProductID = params.ProductID
// 	imageEntity.ImageName = imageName
// 	imageEntity.ImageExtension = extension
// 	imageEntity.ImageSize = int(params.ImageFile.Size)
// 	imageEntity.IsActive = true
// 	return imageEntity, nil
// }
