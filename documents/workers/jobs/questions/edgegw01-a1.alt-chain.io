1. 登录去找麦爷要权限
2. a1 server 上运行着连通kubernetes cluster 的VPN。所以，a1 server充当了 各个集群 master 访问 kubernetes cluster 的入口，这个过程通过ssh做的port-forward