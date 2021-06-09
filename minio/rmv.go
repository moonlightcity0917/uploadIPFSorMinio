package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
)

/*参数样例
opts := minio.RemoveObjectOptions{
GovernanceBypass: true,
VersionID: "myversionid",
}*/
func RmvObject(minioClient *minio.Client, bucket string, opts minio.RemoveObjectOptions) bool {
	err := minioClient.RemoveObject(context.Background(), "mybucket", "myobject", opts)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("RmvObject suc")
	return true
}
