# 一. 概念
## 1. 介绍
1. Xorg是X Window server ，它允许用户通过他们的指尖创造一个图形环境。
2. X.org 项目旨在创建和维护一个自由、可再发行的并且开源的 X11。它是一个开放源码，基于X11桌面所需要的基本软件。
3. Xorg在一个你想运行的硬件和图形软件之间提供了一个接口。除此之外，Xorg也是一个完全的Network-aware，这意味着你可以一个系统上运行一个应用程序,并且还能同时查看其他不同的系统。
# 二. 安装
## 1. 输入驱动程序支持
1. 您将需要通过对您的内核配置更改，来激活对事件接口(CONFIG_INPUT_EVDEV) 的支持。
## 2. 内核模块设置
1. 现代开源的显卡驱动程序依赖于内核模块（KMS）的设置。 KMS提供了改进的图形化引导，更快的用户切换，内置framebuffer终端，便于从控制台Xorg的无缝切换和其他功能。
2. 首先，准备为你的内核配置好KMS。你必须要做这一步,无论你正在使用哪个Xorg显卡驱动程序。
3. 接下来，配置内核使用正确的KMS驱动。Intel，NVIDIA和AMD / ATI是最常见的显卡，所以接下来请按照您自己的显卡来操作。
    - AMDGPU 设置
## 3. 重新构建内核，并重启机器
```
make && make modules_install
mount /boot
make install
grub-mkconfig -o /boot/grub/grub.cfg
reboot
```

# 总
- https://wiki.gentoo.org/wiki/Xorg/Guide/zh-cn