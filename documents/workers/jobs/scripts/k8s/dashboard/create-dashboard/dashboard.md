# 1. 根据 kubernetes 版本选择合适的 kubernetes-dashboard 版本进行安装
1. 官网网址
  - https://github.com/kubernetes/dashboard/releases
# 2. 修改 recommend.yaml
# 3. 创建 ServiceAccount 和 ClusterRoleBinding 
```
vim sa.yaml
```
# 4. 创建 ClusterRoleBinding 
```
vim crb.yaml
```
# 5. 执行
```
kubectl create -f recommend.yaml
kubectl create -f sa.yaml
kubectl create -f crb.yaml
```
# 6. 获取 token
```
kubectl get secret -n kubernetes-dashboard
kubectl describe secret aks-dashboard-admin-token-4qndk -n kubernetes-dashboard
# 把 token 内容粘贴出来以备登录使用
```
# 7. 使用火狐浏览器，谷歌浏览器会校验 ca
```
https://34.222.46.168:31402
```