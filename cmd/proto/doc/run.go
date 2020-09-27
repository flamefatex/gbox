package doc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/flamefatex/config"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) {
	// 获取所以proto文件
	paths, err := getProtoFilePaths(param.Src)
	if err != nil {
		fmt.Println(err)
		return
	}

	docOpts := getTypes(param.Type)

	wg := sync.WaitGroup{}
	for _, opt := range docOpts {
		wg.Add(1)
		go func(opt *docOpt) {
			// 获取命令
			extra := &realCmdbExtra{
				docOpt: opt,
				paths:  paths,
			}
			realCmd := getRealCmd(param, extra)
			//fmt.Println(realCmd.String())

			// 运行
			err = realCmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(opt)
	}
	wg.Wait()

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

// getTypes
type docOpt struct {
	name   string
	suffix string
}

func getTypes(name string) []*docOpt {
	s := []*docOpt{
		{
			name:   "json",
			suffix: "json",
		},
		{
			name:   "markdown",
			suffix: "md",
		},
		{
			name:   "html",
			suffix: "html",
		},
	}
	if name == "all" {
		return s
	}

	for _, o := range s {
		if o.name == name {
			return []*docOpt{o}
		}
	}

	return nil
}

// 命令所需额外信息
type realCmdbExtra struct {
	docOpt *docOpt
	paths  []string
}

// getRealCmd 获取实际运行的命令
func getRealCmd(param *paramInfo, extra *realCmdbExtra) (realCmd *exec.Cmd) {
	// 命令example
	//protoc -Isrc -I/usr/local/include -I$GOPATH/src \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
	//-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	//--doc_out=./doc --doc_opt=markdown,doc.md src/*/*.proto src/*/*/*.proto src/*/*/*/*.proto
	realArgs := []string{
		fmt.Sprintf("-I%s", param.Src),
	}
	imports := config.Config().GetStringSlice("proto.imports")
	for _, ipt := range imports {
		realArgs = append(realArgs, fmt.Sprintf("-I%s", ipt))
	}
	plugins := []string{
		fmt.Sprintf("--doc_out=%s", param.Out),
		fmt.Sprintf("--doc_opt=%s,doc.%s", extra.docOpt.name, extra.docOpt.suffix),
	}
	realArgs = append(realArgs, plugins...)
	realArgs = append(realArgs, extra.paths...)

	realCmd = exec.Command("protoc", realArgs...)
	realCmd.Stdout = os.Stdout
	realCmd.Stderr = os.Stderr
	realCmd.Stdin = os.Stdin

	return
}
