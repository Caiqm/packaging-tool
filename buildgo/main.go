package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

var (
	path   string   // 项目路径
	goarch string   // 系统处理器
	alias  string   // 程序别名
	platform string // 编译平台
)

// 进入目录
func ToDir(mainPath string) error {
	err := os.Chdir(mainPath)
	return err
}

// unix系统打包
func unixBuild(goarch, alias string) (string, error) {
	// cmd := `CGO_ENABLED=0 GOOS=linux GOARCH=` + goarch + ` go build -ldflags="-s -w"`
	cmd := fmt.Sprintf(`CGO_ENABLED=0 GOOS=%s GOARCH=%s  go build -ldflags="-s -w"`, platform, goarch)
	if alias != "" {
		cmd = cmd + " -o " + alias
	}
	fmt.Println("run build linux command:", cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	return strings.TrimSpace(string(result)), err
}

// win系统打包
func winBuild(goarch, alias string) (string, error) {
	cmd := `go build -ldflags="-s -w"`
	if alias != "" {
		cmd = cmd + " -o " + alias
	}
	// 执行的命令
	// cmd = "set CGO_ENABLED=0&&set GOARCH=" + goarch + "&&set GOOS=linux&&" + cmd
	cmd = fmt.Sprintf("set CGO_ENABLED=0&&set GOARCH=%s&&set GOOS=%s&&%s", goarch, platform, cmd)
	fmt.Println("run build linux command:", cmd)
	cmdEc := exec.Command("cmd")
	// 解决cmd双引号问题
	cmdEc.SysProcAttr = &syscall.SysProcAttr{CmdLine: fmt.Sprintf(`/c %s`, cmd), HideWindow: true}
	output, err := cmdEc.CombinedOutput()
	if err != nil {
		return "", errors.New(strings.TrimSpace(string(output)))
	}
	return strings.TrimSpace(string(output)), err
}


func main() {
	dir, _ := os.Getwd()
	flag.StringVar(&path, "path", dir, "Project path. Default current directory")
	flag.StringVar(&goarch, "arch", "amd64", "Computer processor, eg: arm64, arm64")
	flag.StringVar(&alias, "alias", "main", "File alias")
	flag.StringVar(&platform, "platform", "linux", "Compilation platform, eg: linux")
	flag.Parse()
	if path != "" {
		if err := ToDir(path); err != nil {
			fmt.Println("go to the system directory error. err is :", err)
			return
		}
		fmt.Println("cd dir is :", path)
	}
	var err error
	// 判断系统
	if runtime.GOOS == "windows" {
		_, err = winBuild(goarch, alias)
	} else {
		_, err = unixBuild(goarch, alias)
	}
	if err != nil {
		fmt.Println("fail to build, err is :", err)
		return
	}
	// 打包后的名称
	fmt.Println("success build，file path is :", filepath.Join(path, alias))
}