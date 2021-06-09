package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(endPoint string, accesskeyID string, accessKeySecret string, useSSL bool) (*minio.Client, error) {
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskeyID, accessKeySecret, ""),
		Secure: useSSL, //useSSL一般w为false
	})
	if err != nil {
		fmt.Printf("creat minio err %v\n", err)
		return nil, err
	}
	return minioClient, nil
}
