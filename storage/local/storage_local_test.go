package local

import (
	"fmt"
	"github.com/liao0001/cloud_disk/storage"
	"os"
	"testing"
	"time"
)

func TestParseAbs(t *testing.T) {
	dstPath := `/Users/liaoyong/workspace/dev/src/github.com/liao0001/cloud_disk/dst_files`
	conf := storage.ConfigStorage{
		Key:             "local",
		Driver:          "local",
		AccessKeyID:     "",
		AccessKeySecret: "",
		Endpoint:        "/files",
		Bucket:          dstPath,
		WithAcl:         false,
		Expiration:      600,
	}
	serv, err := NewStorage(conf)
	if err != nil {
		panic(err)
	}
	filePath := `/Users/liaoyong/Movies/电视剧/旧版电视剧/尘埃世界/rb/许佳琪.mp4`
	start := time.Now()
	file, _ := os.Open(filePath)
	defer file.Close()

	objName := "test001/测试.mp4"
	link, err := serv.PushFile(objName, file)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now().Sub(start))
	fmt.Println(link)
}
