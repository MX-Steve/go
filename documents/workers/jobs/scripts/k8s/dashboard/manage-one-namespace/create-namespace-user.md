<一个角色控制一个namespace>
1. 在指定命名空间下创建 sa
```
]# kubectl create sa cnych -n kube-system
```
2. 创建 role 和 rolebinding yaml 文件
```
]# vim role.yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: kube-system
  name: role-cnych
rules:
  - apiGroups: ["", "extensions", "apps"]
    resources: ["*"]
    verbs: ["*"]
  - apiGroups: ["batch"]
    resources:
      - jobs
      - cronjobs
    verbs: ["*"]
]# vim role-bind.yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: role-bind-cnych
  namespace: kube-system
subjects:
- kind: ServiceAccount
  name: cnych
  namespace: kube-system
roleRef:
  kind: Role
  name: role-cnych
  apiGroup: rbac.authorization.k8s.io
```
3. 创建 role 和 rolebinding 
```
]# kubect create -f role.yaml
]# kubectl create -f role-bind.yaml
```
4. 找出 token
```
]# kubectl get secret -n kube-system |grep cnych
]# kubectl get secret cnych-token-gxdbv -o jsonpath={.data.token} -n kube-system |base64 -d
```
5. 使用 token 测试登录
    - 登录后要选择你有权限的 namespace 去访问旗下资源
6. 创建 kubeconfig <修改 dashboard-admin.conf>
    - 把所有的 dashboard-admin 替换成 当前用户名
    - 把 token 后面的值改成 当前用户的 token
