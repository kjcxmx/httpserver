#/bin/bash

# httpserver 本地文件服务工具  by kjcxmx@163.com
# GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
# GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm

#变量
TARGET="./target/"
mkdir ${TARGET}

echo "当前golang版本："
go version

echo "编译Unix二进制ing"
GOARCH=386 go build -o ${TARGET}http http.go

echo "编译Mac二进制ing"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${TARGET}http_darwin http.go

echo "编译Windows二进制ing"
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${TARGET}http.exe http.go

echo "编译完成"
tree ${TARGET}
