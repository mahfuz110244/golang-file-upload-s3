package main

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Errors  []FieldError `json:"errors,omitempty"`
	Data    interface{}  `json:"data,omitempty"`
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type FileUploadResponse struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Size      int    `json:"size"`
	Extension string `json:"extension"`
}

type Config struct {
	AwsAccessKeyId       string `env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey   string `env:"AWS_SECRET_ACCESS_KEY"`
	AwsDefaultRegion     string `env:"AWS_DEFAULT_REGION"`
	AwsStorageBucketName string `env:"AWS_STOARGE_BUCKET_NAME"`
}

// Read properties from config.env file
// Command line enviroment variable will overwrite config.env properties
func NewConfig(configFile string) *Config {
	config := Config{}
	err := cleanenv.ReadConfig(configFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	cleanenv.ReadEnv(&config)
	return &config
}

func GenerateFileName(filename string) string {
	t := time.Now()
	_ = t.String()
	currentDatetime := t.Format("02-01-2006-15-04-05PM")
	uid := strings.ToUpper(uuid.New().String()[0:6])
	newfilename := currentDatetime + "-" + uid + "-" + filename
	return newfilename
}

func UploadBulkImage(files []*multipart.FileHeader) (*[]FileUploadResponse, error) {
	conf := NewConfig(".env")
	if conf == nil {
		log.Println("cannot get configuration s3 bukcet")
		return nil, errors.New("custom_error!!!cannot get configuration s3 bucket")
	}
	accessKey := conf.AwsAccessKeyId
	accessSecret := conf.AwsSecretAccessKey
	defaultRegion := conf.AwsDefaultRegion
	bucketName := conf.AwsStorageBucketName
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
		log.Println("aws configuration missing")
		return nil, errors.New("custom_error!!!aws configuration missing")
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
		src, err := file.Open()
		if err != nil {
			log.Printf("failed to open image: %v", err)
		}
		defer src.Close()

		extension := filepath.Ext(file.Filename)
		imageName := file.Filename[:len(file.Filename)-len(extension)]

		key := "files/" + GenerateFileName(file.Filename)

		// Upload the file to S3.
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
			Body:   src,
			// ACL:    aws.String("public-read"),
		})

		//in case it fails to upload
		if err != nil {
			log.Printf("failed to upload image to s3: %v", err)
			// return nil, errors.New("custom_error!!!failed to upload image to s3")
		}
		log.Println("successfully image upload to s3 bucket url: ", result.Location)
		entitys = append(entitys, FileUploadResponse{
			Name:      imageName,
			Url:       result.Location,
			Size:      int(file.Size),
			Extension: extension,
		})
	}
	return &entitys, nil
}

func UploadFileHandler(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	for _, file := range files {
		fmt.Println(file.Filename)
	}
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Files Missing!!! please provide valid files",
		})
	}

	instance, err := UploadBulkImage(files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong!",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Data:    instance,
	})
	// return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", UploadFileHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
