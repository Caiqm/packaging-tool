package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// win系统下执行
func runInWindows(cmd string) (string, error) {
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

// 执行命令
func RunCommand(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

// linux系统下执行
func runInLinux(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), err
}

//根据进程名称获取进程ID
func GetPid(serverName string) (string, error) {
	a := `pidof ` + serverName
	pid, err := RunCommand(a)
	return pid, err
}

// 重新执行main
func RerunMain(pid, mainPath string) (string, error) {
	file, err := os.Stat(mainPath)
	if err != nil {
		return "", err
	}
	var (
		fileMain string
		path     string
		cmd      string
		runCmd   string
	)
	if !file.IsDir() {
		fileMain, path = GetFileAndPath(mainPath)
		err := ToDir(path)
		if err != nil {
			return "", err
		}
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		fmt.Println("cd dir is :", dir)
	} else {
		err := errors.New("must be absolute path")
		return "", err
	}
	if pid != "" {
		cmd = `kill -9 ` + pid
		_, err := RunCommand(cmd)
		if err != nil {
			return "", err
		}
	}
	// runCmd = `nohup ./` + fileMain + ` > nohup.log 2>&1 &`
	runCmd = fmt.Sprintf("nohup ./%s > nohup.log 2>&1 &", fileMain)
	res, err := RunCommand(runCmd)
	return res, err
}

// 获取目录与文件名称
func GetFileAndPath(mainPath string) (string, string) {
	pathSlice := strings.Split(mainPath, "/")
	lastOne := len(pathSlice) - 1
	file := pathSlice[lastOne]
	newPath := strings.Join(pathSlice[:lastOne], "/")
	return file, newPath
}

// 进入目录
func ToDir(mainPath string) error {
	err := os.Chdir(mainPath)
	return err
}

/*
ps ux
USER PID %CPU %MEM VSZ RSS TTY STAT START TIME COMMAND

USER: 进程拥有者
PID:pid
%CPU:占用的cpu使用率
VSZ:占用的内存使用率
RSS:占用的虚拟内存大小
TTY:是否为登入者执行的程序，若为tty1-tty6，为本机登入者，若为pts/??,则为远程登入者。
STAT:程序的状态，R:正在执行中，S：睡眠，T：正在检测或者停止，Z：死亡程序
START:程序开始时间
TIME:程序运行的时间
COMMAND：所执行的指令。
*/
func main() {
	var (
		serverName string // 已启用GO程序名称
		path       string // 启动程序文件路径
		pid        string // 端口号
		showPid    int    // 是否只显示端口号
		err        error  // 错误
	)
	flag.StringVar(&serverName, "p", "", "Server thread")
	flag.StringVar(&path, "m", "", "Golang main file")
	flag.IntVar(&showPid, "pid-only", 0, "show pid")
	flag.Parse()
	// 查询进程id
	if serverName != "" {
		pid, err = GetPid(serverName)
		if err != nil {
			fmt.Println("can not found", serverName)
			fmt.Println(err)
			return
		}
		if pid == "" {
			fmt.Println("main pid is null")
			return
		}
		// 显示端口号
		fmt.Println(pid)
		if showPid == 1 {
			return
		}
	}
	result, err := RerunMain(pid, path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	if result == "" {
		result = "success"
	}
	fmt.Println("result = ", result)
}
