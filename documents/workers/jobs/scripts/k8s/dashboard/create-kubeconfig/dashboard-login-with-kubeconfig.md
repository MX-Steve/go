
# 1. 获取 token 登陆
```
]# kubectl -n kubernetes-dashboard  get  secret  aks-dashboard-admin-token-4qndk
```
- 此时获取的结果可以用于 token 登陆
# 2. 为方便用 Kubeconfig 登陆 dashboard
1. 创建需要的证书
```
]# mkdir -p /usr/local/src/kubernetes/certs
]# cd /usr/local/src/kubernetes
]# openssl genrsa -des3 -passout pass:x -out certs/dashboard.pass.key 2048
]# openssl rsa -passin pass:x -in certs/dashboard.pass.key -out certs/dashboard.key
]# openssl req -new -key certs/dashboard.key -out certs/dashboard.csr -subj '/CN=kube-dashboard'
]# openssl x509 -req -sha256 -days 365 -in certs/dashboard.csr -signkey certs/dashboard.key -out certs/dashboard.crt
]# rm certs/dashboard.pass.key
]# kubectl create secret generic kubernetes-dashboard-certs --from-file=certs -n kubernetes-dashboard
```
2. 获取 token
```
]# DASH_TOCKEN=$(kubectl -n kubernetes-dashboard  get  secret  aks-dashboard-admin-token-4qndk  -o jsonpath={.data.token} |base64 -d)
```
3. 设置 kubeconfig 文件中的一个集群条目
```
]# kubectl config set-cluster kubernetes --server=MASTERIP:8443 --kubeconfig=/usr/local/src/dashbord-admin.conf
```
4. 设置 kubeconfig 文件中的一个用户条目
```
]# kubectl config set-credentials dashboard-admin --token=$DASH_TOCKEN --kubeconfig=/usr/local/src/dashbord-admin.conf
```
5. 设置 kubeconfig 文件中的一个上下文条目
]# kubectl config set-context dashboard-admin@kubernetes --cluster=kubernetes --user=dashboard-admin --kubeconfig=/usr/local/src/dashbord-admin.conf
6. 设置 kubeconfig 文件中的当前上下文
]# kubectl config use-context dashboard-admin@kubernetes --kubeconfig=/usr/local/src/dashbord-admin.conf
]# sz /usr/local/src/dashbord-admin.conf