package tool

import (
	"github.com/gogf/gf/v2/os/genv"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Path 单例selfpath，并导出
var Path = &selfpath{}

type selfpath struct{}

// ExecPath 获取当前可执行文件的路径
func (*selfpath) ExecPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

// ExecRootPath 获取当前项目的根目录
func (s *selfpath) ExecRootPath() (string, error) {
	dir := genv.Get("GF_GCFG_PATH").String()
	if len(dir) > 0 {
		return dir, nil
	}
	dir, err := s.ExecPath()
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(dir, "bin") {
		return path.Dir(dir), nil
	}
	return dir, nil
}

// GetFilePath 获取当前path.go所在路径
func (*selfpath) GetFilePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
