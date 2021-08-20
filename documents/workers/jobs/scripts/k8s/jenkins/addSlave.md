# 1. 背景
1. 已有一台 jenkins-master 正在运行，现在新增 jenkins-slave 节点作为代码构建处理
# 2. jenkins-slave 创建
1. pvc.yaml
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  finalizers:
  - kubernetes.io/pvc-protection
  labels:
    app: jenkins-slave
  name: jenkins-slave-pvc
  namespace: default
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 8Gi
  storageClassName: gp2
  volumeMode: Filesystem
```
2. jenkins-deployment.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
  labels:
    app: jenkins-slave
  name: jenkins-slave
  namespace: default
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: jenkins-slave
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: jenkins-slave
    spec:
      containers:
      - image: jenkins/jenkins:lts
        imagePullPolicy: Always
        name: jenkins-slave
        ports:
        - containerPort: 8080
          protocol: TCP
        - containerPort: 50000
          protocol: TCP
        resources:
          requests:
            cpu: 100m
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /var/jenkins_home
          name: jenkins-home
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext:
        fsGroup: 2000
        runAsNonRoot: true
        runAsUser: 1000
      terminationGracePeriodSeconds: 30
      volumes:
      - name: jenkins-home
        persistentVolumeClaim:
          claimName: jenkins-slave-pvc
```
3. jenkins-service.yaml
```
apiVersion: v1
kind: Service
metadata:
  name: jenkins-slave
  namespace: default
spec:
  externalTrafficPolicy: Cluster
  ports:
  - name: jenkins-www
    nodePort: 30180
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: jenkins-slaves
    nodePort: 30181
    port: 50000
    protocol: TCP
    targetPort: 50000
  selector:
    app: jenkins-slave
  sessionAffinity: None
  type: NodePort
```
4. 执行yaml
```
]# kubectl create -f pvc.yaml
]# kubectl create -f jenkins-deployment.yaml
]# kubectl create -f jenkins-service.yaml
```
# 3. 云平台安全组开放30180 web 端口
# 4. 登录 jenkins-slave 做初始化操作，并初始化用户名密码 admin/admin
# 5. jenkins-master 添加节点
1. https://jenkins.corp.alt-chain.io/computer/ 侧边栏点击 [New Node] 按钮
2. 指定 Node name 名为 jenkinsSlave,选择 Permanent Agent, 点击 OK
3. 设置参数
    1. Name: jenkinsSlave
    2. Description: k8s-jenkins-slave
    3. of executors: 1
    4. Remote root directory: /var/jenkins_home
    5. Labels: k8s-jenkins-slave
    6. Usage: Only build jobs with label expressions matching this node
    7. Launch method: Launch agent by connection it to the master
    8. Internal data directory: remoting
    9. 勾选 Fail if workspace is missing
    10. 点击 Save 保存
4. 保存后提示要下载 agent.jar ,并执行 java 命令
    1. 下载 agent.jar 
    2. 保存 java 命令
# 6. jenkins-slave 上操作
    1. 控制节点把文件传递到容器中
        - kubectl cp agent.jar default/jenkins-slave-56854878c4-5j5zk:/var/jenkins_home/
    2. 进入容器执行命令
        - mkdir /var/jenkins_home/remoting
        - cd /var/jenkins_home
        - nohup java -jar agent.jar -jnlpUrl https://jenkins.corp.alt-chain.io/computer/jenkinsSlave/jenkins-agent.jnlp -secret fd4076d7d187ba2c7bde8dc4c756389acbd097d1166ebc8a6d26ab705f0b9825 -workDir "/var/jenkins_home" -failIfWorkDirIsMissing &
# 7. jenkins-master 上刷新，可以查看到 slave 状态正常，节点添加成功
# 8. 调用 slave 构建项目



