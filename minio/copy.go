package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

/*srcOpts := minio.CopySrcOptions{
Bucket: "my-sourcebucketname",
Object: "my-sourceobjectname",
}
dstOpts := minio.CopyDestOptions{
Bucket: "my-bucketname",
Object: "my-objectname",
}*/
func CopyObject(minioClient *minio.Client, bucket string, dstOpts minio.CopyDestOptions, srcOpts minio.CopySrcOptions) bool {
	uploadInfo, err := minioClient.CopyObject(context.Background(), dstOpts, srcOpts)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Successfully copied object:", uploadInfo)
	return true
}
