package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var rootDir string

func handler(w http.ResponseWriter, r *http.Request) {
	// 打印HTTP请求信息
	fmt.Println("Method:", r.Method, "Host:", r.RemoteAddr, "URL:", r.URL.String())
	fmt.Println("Headers:", r.Header)

	// 调用文件服务器处理器
	fileServer := http.FileServer(http.Dir(rootDir))
	fileServer.ServeHTTP(w, r)
}

// by kjcxmx@163.com

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("https://github.com/kjcxmx/httpserver\tby kjcxmx@163.com\nUsage()\n./http port dir eg:./http 8080")
		return
	}
	// 定义文件根目录
	rootDir = "/"
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("无法获取当前目录，使用根目录：", err)
		if runtime.GOOS == "windows" {
			rootDir = "C:/"
		}
	} else {
		rootDir = dir
	}

	if len(args) == 3 {
		rootDir = args[2]
	}

	// 注册文件服务器处理器到根路径
	http.HandleFunc("/", handler)

	// 定义监听地址和端口号
	addr := ":" + args[1]

	fmt.Printf("https://github.com/kjcxmx/httpserver\tby kjcxmx@163.com\nServer listening on %s\n", addr)
	fmt.Println(http.ListenAndServe(addr, nil))
}
