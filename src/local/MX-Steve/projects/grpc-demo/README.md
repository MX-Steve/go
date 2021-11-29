# go grpc demo interface
# 1. grpc-demo 目录与生成测试idl
``` shell
]# ls 
grpc-demo       idl             proto
# grpc-demo: 项目目录
# idl: idl 配置文件入口
# protoc: 编译后配置文件入口
]# cd idl
]# protoc --go-grpc_out=. schemas/server/djentry.proto
]# protoc --go_out=. schemas/server/djentry.proto
]# mv server ../proto
```
# 2. 编写 grpc server
``` shell
]# vim server.go
]# 
```