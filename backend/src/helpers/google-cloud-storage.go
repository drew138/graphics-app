package helpers

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

// https://adityarama1210.medium.com/simple-golang-api-uploader-using-google-cloud-storage-3d5e45df74a5
// https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-code-sample
var (
	projectID  = os.Getenv("GOOGLE_PROJECT_ID")
	bucketName = os.Getenv("GOOGLE_BUCKET_NAME")
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var Uploader *ClientUploader

func init() {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	Uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "images/",
	}
}

func (c *ClientUploader) UploadToGoogleCloudStorage(file io.Reader, object string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	currentTime := time.Now().String()
	filePath := fmt.Sprintf("%s%s-%s", c.uploadPath, object, strings.Replace(currentTime, " ", "_", -1))
	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(filePath).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://storage.cloud.google.com/%s/images/%s", bucketName, filePath)
	return url, nil
}
