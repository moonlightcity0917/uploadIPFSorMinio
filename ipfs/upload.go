package ipfs

import (
	"bufio"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
)

func UploadIPFS(r io.Reader, ipfsUploadServiceHost string) string {
	sh := shell.NewShell(ipfsUploadServiceHost)
	hash, err := sh.Add(bufio.NewReader(r))
	if err != nil {
		fmt.Println("upload IPFS fail：", err)
		return ""
	}

	fmt.Println("upload IPFS suc：")
	return hash
}
