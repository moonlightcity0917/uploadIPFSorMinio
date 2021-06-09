package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"path"
)

func Upload(objectName string, filePath string, minioClient *minio.Client, bucket string) bool {
	contentType := "application/" + GetFileExt(filePath)
	// Upload
	info, err := minioClient.FPutObject(context.Background(), bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Printf("upload err %v\n", err)
		return false
	}
	fmt.Printf("Success upload minio %s of size %d\n", objectName, info.Size)
	return true
}
func GetFileExt(fileName string) string {
	fileSuffix := path.Ext(fileName)
	fileExt := fileSuffix[1:]
	return fileExt
}
