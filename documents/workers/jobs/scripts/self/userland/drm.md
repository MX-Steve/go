# DRM
## 概念
1. DRM，全称为Direct Rendering Manager。DRM是Linux目前主流的图形显示框架，相比FB架构，DRM更能适应当前日益更新的显示硬件。比如FB原生不支持多层合成，不支持VSYNC，不支持DMA-BUF，不支持异步更新，不支持fence机制等等，而这些功能DRM原生都支持。同时DRM可以统一管理GPU和Display驱动，使得软件架构更为统一，方便管理和维护。
2. DRM是一个内核级的设备驱动，既可以编译到内核中也可以作为标准模块进行加载。DRM最初是在FreeBSD中出现的，后来被移植到Linux系统中，并成为Linux系统的标准部分。
3. DRM从模块上划分，可以简单分为3部分：libdrm、KMS、GEM。
## 组件
### libdrm
1. 对底层接口进行封装，向上层提供通用的API接口，主要是对各种IOCTL接口进行封装。
### KMS
1. Kernel Mode Setting，所谓Mode setting。其实说白了就两件事：更新画面和设置显示参数。
    - 更新画面：显示buffer的切换，多图层的合成方式，以及每个图层的显示位置。
    - 设置显示参数：包括分辨率、刷新率、电源状态（休眠唤醒）等。
2. 包含元素：FB，CRTC，ENCODER，CONNECTOR，PLANE，VBLANK，property。
    1. FB（FrameBuffer）（画布，帧缓冲/帧缓存）
        - 对计算机来说，FrameBuffer 就是一块驱动和应用层都能访问的内存，是唯一一个和硬件无关的基本元素。当然画图之前要有一定的格式化，比方说我可以规定什么样的色彩模式（RGB24 , I420 , YUUV 等等），分辨率是多大等等。
    2. CRTC（绘图现场）
        - 简写翻译过来是阴级摄像管上下文，在 DRM 里 CRTC 就表示显示输出的上下文了。对显示buffer进行扫描，并产生时序信号的硬件模块，通常指Display Controller。首先 CRTC 内指一个 FrameBuffer 地址，外连一个Encoder。
        - 它们俩之间如何沟通？这就是显示模式（ModeSet）要做的事情。ModeSet 包括了像前面提到的色彩模式，还有说显示的时序（timings , ModeLines 等都代表了这个意思）等，通常时序可以按以下来表达：　
        - 一个CRTC 可以连接多个 Encoder ，实现复制屏幕功能。
    3. Encoder（输出转换器）
    4. Connector（连接器 ）
    5. PLANE（硬件图层）
    6. VBLANK （垂直消隐，场消隐）
    7. property （属性）
### GEM
1. Graphic Execution Manager，主要负责显示buffer的分配和释放，也是GPU唯一用到DRM的地方。
2. 包含元素：DUMB，PRIME，fence。
    1. DUMB
        - 只支持连续物理内存，基于kernel中通用CMA API实现，多用于小分辨率简单场景。
    2. PRIME
        - 连续、非连续物理内存都支持，基于DMA-BUF机制，可以实现buffer共享，多用于大内存复杂场景。
    3. fence
        - buffer同步机制，基于内核dma_fence机制实现，用于防止显示内容出现异步问题。
# DRM支持DRI的方式
1. DRM提供到显卡硬件的同步访问。
    - Direct rendering system有多个实体（比如X server，多个direct-rendering客户端，以及kernel）竞争访问显卡硬件。PC类的显卡在多个实体访问显卡硬件时会使用锁。DRM为每个显卡设备提供了一个锁，来同步硬件的访问。比如X server正在执行2D渲染，此direct-rendering客户端执行一个软件回调，这个软件回调会读写frame buffer。对于一些高端卡来说，由于硬件内部本身会对访问命令做排序，因此并不需要使用这个锁。
2. DRM在访问显卡硬件时，强制执行DRI安全测策略。
    - X server以root权限运行，在访问显卡的framebuffer和MMIO区域时，会用/dev/mem映射这些区域。direct-rendering 客户端，并不是运行在root权限的，但是仍然需要类似的映射。DRM设备接口允许客户端创建这些映射，但是必须遵守以下限制：
        - * 仅当客户端连接到X server时才能映射这些区域，这就迫使direct-rendering客户端遵守正常的X server安全策略。 
        - * 仅当客户端能够打开/dev/drm?时才可以映射这些区域。这允许系统管理员可以配置direct rendering访问，仅可信的用户才能访问。 
        - * 客户端只能映射X server允许映射的区域。
3. DRM提供了一个通用的DMA引擎。
    - 大部分现代PC类计算机的显卡硬件提供command FIFO的DMA访问。
    - DMA 访问比MMIO访问有更好的吞吐量性能。对于这些显卡，DRM 提供的DMA引擎包含下面的features: * 