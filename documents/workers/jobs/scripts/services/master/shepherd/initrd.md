解压：

]# xz -d initrd-superdata-5.3.0-62-generic.img.xz 
]# ls 
initrd-superdata-5.3.0-62-generic.img
]# cpio -idv < initrd-superdata-5.3.0-62-generic.img        
]# ls 
bin  etc   init                                   lib   root  tmp  var
dev  home  initrd-superdata-5.3.0-62-generic.img  proc  sys   usr

压缩：
gen_initramfs_list.sh ./Filesystem/ >filelist  
gen_init_cpio filelist >rootfs.cpio
xz rootfs.cpio
mv rootfs.cpio.xz initrd-superdata-tmp.img.xz 