package ipfs

import (
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

func CatIPFS(cid string, ipfsUploadServiceiHost string) ([]byte, error) {
	sh := shell.NewShell(ipfsUploadServiceiHost)
	read, err := sh.Cat(cid)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(read)
	return body, nil
}
