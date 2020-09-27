package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)

// 内容的read buffer，默认64KB
const readBufferSize = 64 * 1024

var (
	readBuffer sync.Pool
)

func newReadBufferPool() []byte {
	if br := readBuffer.Get(); br != nil {
		return br.([]byte)
	}
	return make([]byte, readBufferSize)
}

func putReadBufferPool(br []byte) {
	readBuffer.Put(br)
}

func Md5(filePath string) (md5Str string, err error) {
	// 缓存区
	buf := newReadBufferPool()
	defer func() {
		putReadBufferPool(buf)
	}()

	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	//关闭文件
	defer f.Close()

	md5Hash := md5.New()
	_, err = io.CopyBuffer(md5Hash, f, buf)
	if err != nil {
		return
	}
	md5Str = fmt.Sprintf("%x", md5Hash.Sum(nil))
	return
}
