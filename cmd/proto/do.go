package proto

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/flamefatex/config"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

func Do(cmd *cobra.Command, args []string) {
	// 获取所以proto文件
	paths, err := getProtoFilePaths(param.Src)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 建临时目录
	tmpDir, err := ioutil.TempDir("", "gbox_proto_*")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 清理临时目录
	defer os.RemoveAll(tmpDir)

	wg := sync.WaitGroup{}
	for _, p := range paths {
		wg.Add(1)
		go func(p string) {
			// 获取参数
			baseArgs := getBaseArgs(param, tmpDir)
			realArgs := append(baseArgs, p)

			// 执行
			realCmd := exec.Command("protoc", realArgs...)
			realCmd.Stdout = os.Stdout
			realCmd.Stderr = os.Stderr
			realCmd.Stdin = os.Stdin
			//fmt.Println(realCmd.String())
			err = realCmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()

	// 复制
	err = copy.Copy(tmpDir+param.PackageRoot, param.Out)
	if err != nil {
		fmt.Println(err)
	}

}

// getProtoFilePaths
func getProtoFilePaths(src string) (paths []string, err error) {
	paths = make([]string, 0)
	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // ignore errors walking in file system
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".proto") {
			return nil
		}
		paths = append(paths, path)
		return nil
	})
	return
}

func getBaseArgs(param *Param, tmpDir string) (baseArgs []string) {
	// 命令example
	//protoc -Isrc -I/usr/local/include -I$GOPATH/src \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	//--go_out=plugins=grpc:/tmp/protos \
	//--govalidators_out=/tmp/protos \
	//--grpc-gateway_out=logtostderr=true:/tmp/protos \
	//$f
	baseArgs = []string{
		fmt.Sprintf("-I%s", param.Src),
		//fmt.Sprintf("-I%s", "/usr/local/include"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src/github.com/grpc-ecosystem/grpc-gateway"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"),
		fmt.Sprintf("--go_out=plugins=grpc:%s", tmpDir),
		fmt.Sprintf("--grpc-gateway_out=logtostderr=true:%s", tmpDir),
		fmt.Sprintf("--govalidators_out=%s", tmpDir),
	}

	imports := config.Config().GetStringSlice("proto.imports")
	for _, ipt := range imports {
		baseArgs = append(baseArgs, fmt.Sprintf("-I%s", ipt))
	}

	return
}
