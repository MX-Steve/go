1. 按期代理
```shell
go env -w GOPROXY=https://goproxy.io,direct
go env -w GO111MODULE=on
```
2. 安装 gopls 
```shell
go get golang.org/x/tools/gopls
```
3. 关闭代理
```shell
go env -w GO111MODULE=off
```