package storage

import (
	"FileStore-Server/config"
	"bufio"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"path"
)

var sh *shell.Shell

//数据上传到ipfs
func UploadIPFS(r io.Reader) string {
	sh = shell.NewShell(config.IpfsUploadServiceHost)
	hash, err := sh.Add(bufio.NewReader(r))
	if err != nil {
		fmt.Println("upload IPFS fail：", err)
		return ""
	}
	fmt.Println("upload IPFS suc：")
	return hash
}

//从ipfs下载数据
func CatIPFS(cid string) ([]byte, error) {
	sh = shell.NewShell(config.IpfsUploadServiceHost)
	read, err := sh.Cat(cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(read)

	return body, nil
}

//对象存储
func UploadMinio(objectName string, filePath string, endPoint string, accesskeyID string, accessKeySecret string, bucket string) bool {
	fmt.Println("enter UploadMinio")
	ctx := context.Background()
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskeyID, accessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Printf("creat minio err %v\n", err)
		return false
	}
	contentType := "application/" + GetFileExt(filePath)
	// Upload
	info, err := minioClient.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Printf("upload err %v\n", err)
		return false
	}
	fmt.Printf("Success upload minio %s of size %d\n", objectName, info.Size)
	return true
}

func DownloadMinio(objectName string, filePath string, endPoint string, accesskeyID string, accessKeySecret string, bucket string) bool {
	ctx := context.Background()
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accesskeyID, accessKeySecret, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Printf("creat minio err %v\n", err)
		return false
	}

	err = minioClient.FGetObject(ctx, bucket, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Printf("DownloadMinio err %v\n", err)
		return false
	}
	return true
}

func GetFileExt(fileName string) string {
	fileSuffix := path.Ext(fileName)
	fileExt := fileSuffix[1:]
	return fileExt
}