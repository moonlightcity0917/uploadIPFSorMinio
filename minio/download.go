package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"os"
)

func DownloadFile(objectName string, filePath string, minioClient *minio.Client, bucket string) bool {
	err := minioClient.FGetObject(context.Background(), bucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("DownloadMinio err %v\n", err)
		return false
	}
	fmt.Printf("DownloadFile suc")
	return true
}

func DownloadFolw(objectName string, filePath string, minioClient *minio.Client, bucket string) bool {
	object, err := minioClient.GetObject(context.Background(), bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	localFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("DownloadFolw suc")
	return true
}
