# 1. 现象
1. 执行 kubectl top nodes 正常
```
[xxx]# kubectl top nodes
NAME                                          CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%     
ip-xxx   1942m        97%    5780Mi          38%         
ip-xxx    144m         7%     4553Mi          60%         
ip-xxx   2763m        69%    8630Mi          57%         
ip-xxx   195m         9%     1527Mi          40%    
```
2. 执行 kubectl top pods 不正常
```
[xxx]# kubectl top pods -n monitoring
W0714 03:41:48.359570   22248 top_pod.go:274] Metrics not available for pod monitoring/alertmanager-main-0, age: 503h57m41.359559229s
error: Metrics not available for pod monitoring/alertmanager-main-0, age: 503h57m41.359559229s
```
# 2. 原因
1. github 上似乎找到原因了，原来是我的docker版本是 18 版本的，对 metrics 模块有影响，是 bug，你可以将 docker 升级到 19 或以上，要不等着 k8s 修复这个 bug  
