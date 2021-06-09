package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

/*opts:=minio.ListObjectsOptions{
Prefix: "myprefix",
Recursive: true,
}*/
func ListObject(minioClient *minio.Client, bucket string, opts minio.ListObjectsOptions) minio.ObjectInfo {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucket, opts)
	var object minio.ObjectInfo
	for object = range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return object
		}
		fmt.Println(object)
		return object
	}
	return object
}
