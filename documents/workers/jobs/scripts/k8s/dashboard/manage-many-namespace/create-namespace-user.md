<一个角色控制多个namespace>   
# 1. 创建集群级别的角色 ClusterRole
```
]# vim ClusterRole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: control-miner
rules:
  - apiGroups:
    - '*'
    resources:
    - '*'
    verbs:
    - '*'
  - nonResourceURLs:
    - '*'
    verbs:
    - '*'
]# kubectl create -f ClusterRole.yaml
```
# 2. 在 default 命名空间创建 ServiceAccount
```
]# kubectl create serviceaccount miner
```
# 3. 对sa和集群角色建立绑定关系
```
]# kubectl create rolebinding control-miner-binding --clusterrole=control-miner --serviceaccount=default:miner --namespace=default
]# kubectl create rolebinding control-miner-binding --clusterrole=control-miner --serviceaccount=default:miner --namespace=data
]# kubectl create rolebinding control-miner-binding --clusterrole=control-miner --serviceaccount=default:miner --namespace=bminer
```
- 这样不便于统一管理，则把 rolebinding 落文件
```
]# vim rolebinding.yaml 
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: control-miner-binding
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: control-miner
subjects:
- kind: ServiceAccount
  name: miner
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: control-miner-binding
  namespace: bminer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: control-miner
subjects:
- kind: ServiceAccount
  name: miner
  namespace: default
```
# 4. 获取token先测试一下
```
]# token_name=`kubectl get secret  -n default |grep miner-token |awk '{print $1}'`
]# secret=`kubectl get secret $token_name -o jsonpath={.data.token} -n default |base64 -d`
]# echo $secret 
```
# 5. 生成config文件
```
]# kubectl config set-cluster kubernetes --server=172.20.58.129:8443 --kubeconfig=./dashboard-miner.conf
# 这里的scret参数需要替换成上面获取到的登陆的token值
]# kubectl config set-credentials dashboard-miner --token="$secret" --kubeconfig=./dashboard-miner.conf
]# kubectl config set-context dashboard-miner@kubernetes --cluster=kubernetes --user=dashboard-miner --kubeconfig=./dashboard-miner.conf
]# kubectl config use-context dashboard-miner@kubernetes  --kubeconfig=./dashboard-miner.conf
```
# 6. 导出 config 文件测试