## 1. 概念
1. 是一种用于内核调试的虚拟文件系统，内核开发者通过 debugfs 和用户空间交换数据
2. 不实际存储在硬盘上，而是 Linux 内核运行起来后才建立起来
## 2. 使用
```shell
mount -t debugfs none /your/debugfs/dir
```
