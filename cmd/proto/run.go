package proto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/flamefatex/config"
	"github.com/flamefatex/gbox/pkg/md5"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	// 获取所以proto文件
	paths, err := getProtoFilePaths(param.Src)
	if err != nil {
		fmt.Printf("get proto files err:%s\n", err)
		return
	}

	// 建临时目录
	tmpDir, err := ioutil.TempDir("", "gbox_proto_*")
	if err != nil {
		fmt.Printf("make temp dir err:%s\n", err)
		return
	}
	// 清理临时目录
	defer os.RemoveAll(tmpDir)

	// 获取
	oldMd5Map := getMetaData()
	newMd5Map := make(map[string]string)

	hasChanged := false

	wg := sync.WaitGroup{}
	for _, p := range paths {
		md5Str, err := md5.Md5(p)
		if err != nil {
			fmt.Printf("md5 file:%s,err:%s\n", p, err)
			continue
		}
		newMd5Map[p] = md5Str
		// 没有改变跳过
		if oldMd5Str, ok := oldMd5Map[p]; ok && oldMd5Str == md5Str {
			continue
		}
		hasChanged = true

		wg.Add(1)
		go func(p string) {
			// 获取命令
			extra := &realCmdbExtra{
				tmpDir:   tmpDir,
				filePath: p,
			}
			realCmd := getRealCmd(param, extra)
			//fmt.Println(realCmd.String())

			// 运行
			err = realCmd.Run()
			if err != nil {
				fmt.Printf("run cmd err:%s\n", err)
			}
			wg.Done()
		}(p)
	}
	wg.Wait()

	if !hasChanged {
		return
	}

	saveMetaData(newMd5Map)

	// 复制
	err = copy.Copy(tmpDir+param.PackageRoot, param.Out)
	if err != nil {
		fmt.Printf("copy err:%s\n", err)
	}

}

// getProtoFilePaths 获取proto 文件
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

// 命令所需额外信息
type realCmdbExtra struct {
	tmpDir   string // 临时目录路径
	filePath string // proto文件路径
}

// getRealCmd 获取实际运行的命令
func getRealCmd(param *paramInfo, extra *realCmdbExtra) (realCmd *exec.Cmd) {
	// 命令example
	//protoc -Isrc -I/usr/local/include -I$GOPATH/src \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	//--go_out=plugins=grpc:/tmp/protos \
	//--govalidators_out=/tmp/protos \
	//--grpc-gateway_out=logtostderr=true:/tmp/protos \
	//$f
	realArgs := []string{
		fmt.Sprintf("-I%s", param.Src),
		//fmt.Sprintf("-I%s", "/usr/local/include"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src/github.com/grpc-ecosystem/grpc-gateway"),
		//fmt.Sprintf("-I%s", "/Users/flame/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"),

	}
	imports := config.Config().GetStringSlice("proto.imports")
	for _, ipt := range imports {
		realArgs = append(realArgs, fmt.Sprintf("-I%s", ipt))
	}
	plugins := []string{
		fmt.Sprintf("--go_out=plugins=grpc:%s", extra.tmpDir),
		fmt.Sprintf("--grpc-gateway_out=logtostderr=true:%s", extra.tmpDir),
		fmt.Sprintf("--govalidators_out=%s", extra.tmpDir),
	}
	realArgs = append(realArgs, plugins...)
	realArgs = append(realArgs, extra.filePath)

	realCmd = exec.Command("protoc", realArgs...)
	realCmd.Stdout = os.Stdout
	realCmd.Stderr = os.Stderr
	realCmd.Stdin = os.Stdin

	return
}

func getMetaData() (md5Map map[string]string) {
	// automock执行前的原始的meta文件
	md5Map = make(map[string]string)

	b, err := ioutil.ReadFile(".gbox_proto")
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("ReadFile .gbox_proto err:%s\n", err)
		return
	}
	if len(b) != 0 {
		err = json.Unmarshal(b, &md5Map)
		if err != nil {
			fmt.Printf("json.Unmarshal .gbox_proto err:%s\n", err)
			return
		}
	}
	return
}

func saveMetaData(md5Map map[string]string) {
	b, err := json.Marshal(md5Map)
	if err != nil {
		fmt.Printf("json.Marshal .gbox_proto err:%s\n", err)
		return
	}

	err = ioutil.WriteFile(".gbox_proto", b, os.ModePerm)
	if err != nil {
		fmt.Printf("WriteFile .gbox_proto err:%s\n", err)
		return
	}
	return

}
