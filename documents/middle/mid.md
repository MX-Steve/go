# 1. 生成跨平台的二进制包
```bash
$ go get github.com/mitchellh/gox
$ cd $GOPATH/src/github.com/mitchellh/gox
$ go build
# cp gox.exe 到 path 路径
# 进入到项目目录
$ gox
# 会生成所有平台二进制文件
$ gox -os "darwin linux"
# 生成 mac 平台文件
$ gox -os "linux/amd64"
# 生成 linux 平台文件
```
# 2. git 忽略告警
```bash
# warning: LF will be replaced by CRLF in src/local/go.mod.
$ git config core.autocrlf false
$ git config --get core.autocrlf
```
