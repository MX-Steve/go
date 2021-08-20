# 1. Dockerfile 镜像
1. ubuntu:bionic 镜像为基础
2. DEBIAN_FRONTEND=noninteractive
    - 直接运行命令，而无需向用户请求输入（所有操作都是非交互式的）
    - 请确保只在Dockerfile中调用的RUN命令中设置了该选项，而不是使用ENV命令进行全局的设置。
    - 因为ENV命令在整个容器运行过程中都会生效，所以当你通过BASH和容器进行交互时，如果进行了全局设置那就会出问题。
3. 安装包的介绍
    - build-essential
        - 安装了该软件包，编译c/c++所需要的软件包也都会被安装。因此如果想在Ubuntu中编译c/c++程序,只需要安装该软件包就可以了。
    - libelf-dev
        - libelf 的开发包
        - libelf 操作elf文件的开发库
    - linux-headers-$KERNEL_VERSION
        - 更新 kernel headers
    - dkms
        - 帮我们维护内核外的这些驱动程序，在内核版本变动之后可以自动重新生成新的模块。
    - kexec-tools
        - 生成kexec工具
        - 内核调试，Kexec 大大减少了重启引起的系统宕机时间
    - wget       
        - 下载工具 
4. 下载包的介绍
    - bminer-runtime 
        - bminer 运行工具
    - linux-modules-$KERNEL_VERSION linux-modules-extra-$KERNEL_VERSION
        - 下载指定版本
5. amd驱动下载
    - amdgpu-pro-$AMDGPU_PRO_DRIVER_VERSION        
    - dpkg -i xxx.deb
        - 安装驱动软件包
6. mkdir -p /out/etc /out/lib/firmware/$KERNEL_VERSION /out/lib/modules/$KERNEL_VERSION/updates/ /out/sbin
    - 创建相应目录存放驱动文件等  
    - cp -ar /opt /out/
7. cp -ar /opt /out/ 
8. cp -ar /sbin/kexec /out/sbin/ 
9. cp -ar /lib/modules/$KERNEL_VERSION/updates/dkms /out/lib/modules/$KERNEL_VERSION/updates/ 
10. ar -p /tmp/bminer-runtime*.deb data.tar.xz|tar Jx -C /out      
    - deb 包组成
        - control.tar.xz  
            - control
                - 包的元数据信息
        - data.tar.xz  
            - etc  lib  lib64  usr 等系统文件
        - debian-binary
            - 2.0 : deb 版本
    - 将 bminer-runtime*.deb 中 data.tar.xz 数据解压到 /out 目录下
11. ar -p /tmp/linux-modules-$KERNEL_VERSION*.deb data.tar.xz|tar Jx -C /out \ 
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/gpu/drm/drm.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/gpu/drm/drm_kms_helper.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/i2c/algos/i2c-algo-bit.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/video/fbdev/core/fb_sys_fops.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/video/fbdev/core/syscopyarea.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/video/fbdev/core/sysfillrect.ko \
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/video/fbdev/core/sysimgblt.ko
    - 将 data.tar.xz 中 lib/modules/里的指定文件拷贝到 /out 中
12. ar -p /tmp/linux-modules-extra-*.deb data.tar.xz|tar Jx -C /out \ 
    ./lib/modules/$KERNEL_VERSION/kernel/drivers/iommu/amd_iommu_v2.ko
13. cp -ar /etc/OpenCL /etc/ld.so.conf /etc/ld.so.conf.d /etc/ld.so.cache /out/etc/
14. cp -ar /usr/src/amdgpu-*/firmware/amdgpu /out/lib/firmware/$KERNEL_VERSION/ 
15. FROM busybox:latest
    COPY --from=builder /out/ /       
    - 启动第二个镜像 busybox，将 /out 目录覆盖到 busybox / 下

# 2. build.sh 脚本
1. docker build --force-rm -t shepherd-userland-amdgpu-pro:4.15.0-50-generic --build-arg=KERNEL_VERSION="4.15.0-50-generic" amdgpu-pro
    - 使用 amdgpu-pro 目录下的 Dockerfile 制作进行
    - --build-arg=KERNEL_VERSION="4.15.0-50-generic"
        - 给 Dockerfile 里传递参数 KERNEL_VERSION
    - --force-rm 
        - 制作完强制删除
    - -t shepherd-userland-amdgpu-pro:4.15.0-50-generic
        - 给制作的镜像起名
2. docker export $DOCKER_INSTANCE|bsdtar -acf crater-userland-5.4.0-71-generic.tar.xz --exclude dev/console @-
    - 将 busybox 实例的文件系统导出成 crater-userland-5.4.0-71-generic.tar.xz 包        